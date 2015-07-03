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

// Interceptor contains functions that receive (and can modify) a
// wrecker.Request before it is sent to the server.  The Wrecker instance
// maintains an array of Interceptors that are applied to every
// wrecker.Request in the order that they were assigned.
type Interceptor struct {
	Request     func(*Request) error
	RawRequest  func(*http.Request) error
	RawResponse func(*http.Response) error
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

func (w *Wrecker) newRequest(httpVerb string, endpoint string) *Request {
	return &Request{
		HttpVerb:      httpVerb,
		Endpoint:      endpoint,
		Params:        url.Values{},
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

func (w *Wrecker) Delete(endpoint string) *Request {
	return w.newRequest(DELETE, endpoint)
}

// Interceptor adds a new InterceptorFunc into the array of
// functions that are applied to each wrecker.Request *before* it is sent
// to the server.
func (w *Wrecker) Interceptor(interceptor Interceptor) {
	w.Interceptors = append(w.Interceptors, interceptor)
}

func (w *Wrecker) sendRequest(r *Request) (*http.Response, error) {
	var contentType string
	var bodyReader io.Reader
	var err error

	// Apply Requesst Interceptors
	for _, interceptor := range w.Interceptors {
		if interceptor.Request != nil {
			if err := interceptor.Request(r); err != nil {
				return nil, err
			}
		}
	}

	// Empty Body means that we're posting Params via Form encoding
	if r.HttpBody == nil {
		bodyReader = strings.NewReader(r.Params.Encode())
		contentType = w.DefaultContentType
	} else {
		// Otherwise, we're sending a request body
		contentType = "application/json"
		bodyReader, err = prepareRequestBody(r.HttpBody)

		if err != nil {
			return nil, err
		}
	}

	// Create the HTTP client request
	clientReq, err := http.NewRequest(r.HttpVerb, r.URL(), bodyReader)
	if err != nil {
		return nil, err
	}

	clientReq.Header.Add("Content-Type", contentType)

	// Add headers to clientReq
	for key, value := range r.Headers {
		clientReq.Header.Add(key, value)
	}

	// Apply Requesst Interceptors
	for _, interceptor := range w.Interceptors {
		if interceptor.RawRequest != nil {
			if err := interceptor.RawRequest(clientReq); err != nil {
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

	// Apply RawResponse Interceptors
	for _, interceptor := range w.Interceptors {
		if interceptor.RawResponse != nil {
			if err := interceptor.RawResponse(resp); err != nil {
				return nil, err
			}
		}
	}

	// Packing into response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, r.Response)
	return resp, err
}

// prepareRequestBody() function was originally included in the
// github.com/franela/goreq application (which is also MIT licensed)
func prepareRequestBody(b interface{}) (io.Reader, error) {

	// try to jsonify it
	j, err := json.Marshal(b)

	if err == nil {
		return bytes.NewReader(j), nil
	}
	return nil, err
}
