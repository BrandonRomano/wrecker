package wrecker

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Wrecker struct {
	BaseURL            string
	HttpClient         *http.Client
	DefaultContentType string
}

func New(baseUrl string) *Wrecker {
	return &Wrecker{
		BaseURL: baseUrl,
		HttpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		DefaultContentType: "application/x-www-form-urlencoded",
	}
}

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
)

func (w *Wrecker) newRequest(httpVerb string, endpoint string) *WreckerRequest {
	return &WreckerRequest{
		HttpVerb:      httpVerb,
		Endpoint:      endpoint,
		Params:        url.Values{},
		Headers:       make(map[string]string),
		WreckerClient: w,
	}
}

func (w *Wrecker) Get(endpoint string) *WreckerRequest {
	return w.newRequest(GET, endpoint)
}

func (w *Wrecker) Post(endpoint string) *WreckerRequest {
	return w.newRequest(POST, endpoint)
}

func (w *Wrecker) Put(endpoint string) *WreckerRequest {
	return w.newRequest(PUT, endpoint)
}

func (w *Wrecker) Delete(endpoint string) *WreckerRequest {
	return w.newRequest(DELETE, endpoint)
}

func (w *Wrecker) sendRequest(r *WreckerRequest) error {

	var contentType string
	var bodyReader io.Reader
	var err error

	// Empty Body means that we're posting Params via Form encoding
	if r.Body == nil {

		bodyReader = strings.NewReader(r.Params.Encode())
		contentType = w.DefaultContentType

	} else {

		// Otherwise, we're sending a request body
		if bodyReader, err = prepareRequestBody(r.Body); err != nil {
			return err
		}

		contentType = "application/json"
	}

	// Create the HTTP client request
	clientReq, err := http.NewRequest(r.HttpVerb, r.URL(), bodyReader)
	if err != nil {
		return err
	}

	clientReq.Header.Add("Content-Type", contentType)

	// Add headers to clientReq
	for key, value := range r.Headers {
		clientReq.Header.Add(key, value)
	}

	// Executing request
	clientRes, err := w.HttpClient.Do(clientReq)
	if err != nil {
		return err
	}
	defer clientRes.Body.Close()

	// Packing into response
	body, err := ioutil.ReadAll(clientRes.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, r.Response)
}

// prepareRequestBody() function was originally included in the
// github.com/franela/goreq application (which is also MIT licensed)
func prepareRequestBody(b interface{}) (io.Reader, error) {
	switch b.(type) {
	case string:
		// treat is as text
		return strings.NewReader(b.(string)), nil
	case io.Reader:
		// treat is as text
		return b.(io.Reader), nil
	case []byte:
		//treat as byte array
		return bytes.NewReader(b.([]byte)), nil
	case nil:
		return nil, nil
	default:
		// try to jsonify it
		j, err := json.Marshal(b)
		if err == nil {
			return bytes.NewReader(j), nil
		}
		return nil, err
	}
}
