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
	RequestInterceptor func(*Request) error
}

func New(baseUrl string) *Wrecker {
	return &Wrecker{
		BaseURL: baseUrl,
		HttpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		DefaultContentType: "application/x-www-form-urlencoded",
		RequestInterceptor: nil,
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

func (w *Wrecker) sendRequest(r *Request) (*http.Response, error) {

	var contentType string
	var bodyReader io.Reader
	var err error

	// GET methods don't have an HTTP Body.  For all other methods,
	// it's time to defined the body content.
	if r.HttpVerb != GET {

		bodyReader, contentType, err = w.getRequestBody(r)

		if err != nil {
			return nil, err
		}
	}

	// Create the HTTP client request
	clientReq, err := http.NewRequest(r.HttpVerb, r.URL(), bodyReader)
	if err != nil {
		return nil, err
	}

	// Add Basic Auth, if we have it
	if r.BasicAuthInfo != nil {
		clientReq.SetBasicAuth(r.BasicAuthInfo.Username, r.BasicAuthInfo.Password)
	}

	// Set Content-Type for this request
	clientReq.Header.Add("Content-Type", contentType)

	// Add headers to clientReq
	for key, value := range r.Headers {
		clientReq.Header.Add(key, value)
	}

	// Add cookies
	for _, cookie := range r.Cookies {
		clientReq.AddCookie(cookie)
	}

	// Executing request
	resp, err := w.HttpClient.Do(clientReq)
	if err != nil {
		return nil, err
	}

	// Packing into response, if we have one
	if r.Response != nil {
		body, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(body, r.Response)
		if err != nil {
			return resp, err
		}
	}

	// Handling HTTP Error
	if !(resp.StatusCode >= 200 && resp.StatusCode < 300) {
		return resp, ResponseError{
			StatusCode: resp.StatusCode,
			StatusText: resp.Status,
		}
	}

	// OK
	return resp, nil
}

func (w *Wrecker) getRequestBody(r *Request) (io.Reader, string, error) {

	// If there's no body, then try form encoding
	if r.HttpBody == nil {

		// If there are Form Parameters, then let's use form
		return strings.NewReader(r.FormParams.Encode()), "application/x-www-form-urlencoded", nil

	} else {

		// Otherwise, try using a JSON encoded body that was given to us

		var reader io.Reader
		var contentType = "application/json"
		var err error

		switch r.HttpBody.(type) {

		case io.Reader:
			reader = r.HttpBody.(io.Reader)

		case []byte:
			reader = bytes.NewReader(r.HttpBody.([]byte))

		default:

			// try to jsonify it
			j, err := json.Marshal(r.HttpBody)

			if err == nil {
				reader = bytes.NewReader(j)
			}
		}

		return reader, contentType, err
	}

}
