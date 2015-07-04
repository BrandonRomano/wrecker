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
	Interceptors       []Interceptor
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
	PATCH  = "PATCH"
	DELETE = "DELETE"
)

func (w *Wrecker) newRequest(httpVerb string, endpoint string) *Request {
	return &Request{
		HttpVerb:      httpVerb,
		Endpoint:      endpoint,
		URLParams:     url.Values{},
		FormParams:    url.Values{},
		Headers:       make(map[string]string),
		WreckerClient: w,
	}
}

func (w *Wrecker) Get(endpoint string) *Request {
	return w.newRequest(GET, endpoint)
}

func (w *Wrecker) Post(endpoint string) *Request {
	return w.newRequest(POST, endpoint)
}

func (w *Wrecker) Put(endpoint string) *Request {
	return w.newRequest(PUT, endpoint)
}

func (w *Wrecker) Patch(endpoint string) *Request {
	return w.newRequest(PATCH, endpoint)
}

func (w *Wrecker) Delete(endpoint string) *Request {
	return w.newRequest(DELETE, endpoint)
}

// Interceptor adds a new InterceptorFunc into the array of
// functions that are applied to each wrecker.Request *before* it is sent
// to the server.
func (w *Wrecker) Intercept(interceptor Interceptor) *Wrecker {
	w.Interceptors = append(w.Interceptors, interceptor)

	return w
}

func (w *Wrecker) sendRequest(r *Request) (*http.Response, error) {

	var contentType string = "application/x-www-form-urlencoded"
	var bodyReader io.Reader
	var err error

	// Apply WreckerRequest Interceptors
	for _, interceptor := range w.Interceptors {
		if interceptor.WreckerRequest != nil {
			if err := interceptor.WreckerRequest(r); err != nil {
				return nil, err
			}
		}
	}

	// GET methods don't have an HTTP Body.  For all other methods,
	// it's time to defined the body content.
	if r.HttpVerb != GET {

		if r.HttpBody != nil {

			// Otherwise, try using a JSON encoded body that was given to us
			contentType = "application/json"

			bodyReader, err = prepareRequestBody(r.HttpBody)

		} else {

			// If there are Form Parameters, then let's use form
			bodyReader = strings.NewReader(r.FormParams.Encode())
		}
	}

	// Create the HTTP client request
	clientReq, err := http.NewRequest(r.HttpVerb, r.URL(), bodyReader)
	if err != nil {
		return nil, err
	}

	// Set Content-Type for this request
	clientReq.Header.Add("Content-Type", contentType)

	// Add headers to clientReq
	for key, value := range r.Headers {
		clientReq.Header.Add(key, value)
	}

	// Apply Requesst Interceptors
	for _, interceptor := range w.Interceptors {
		if interceptor.HTTPRequest != nil {
			if err := interceptor.HTTPRequest(clientReq); err != nil {
				return nil, err
			}
		}
	}

	// Executing request
	resp, err := w.HttpClient.Do(clientReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Packing into response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Apply HTTPResponse Interceptors
	for _, interceptor := range w.Interceptors {
		if interceptor.HTTPResponse != nil {
			if err := interceptor.HTTPResponse(resp, body); err != nil {
				return nil, err
			}
		}
	}

	err = json.Unmarshal(body, r.Response)
	return resp, err
}

func prepareRequestBody(b interface{}) (io.Reader, error) {

	switch b.(type) {

	case io.Reader:
		return b.(io.Reader), nil

	case []byte:
		return bytes.NewReader(b.([]byte)), nil

	default:

		// try to jsonify it
		j, err := json.Marshal(b)

		if err == nil {
			return bytes.NewReader(j), nil
		}
		return nil, err
	}
}
