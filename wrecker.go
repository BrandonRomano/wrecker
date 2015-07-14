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

type Interceptor struct {
	WreckerRequest func(*Request) error
}

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

type Person struct {
	Value string
}

func NewPerson(value string) *Person {
	return &Person{
		Value: value,
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

	var contentType string
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

	// Set Content-Type for this request
	clientReq.Header.Add("Content-Type", contentType)

	// Add headers to clientReq
	for key, value := range r.Headers {
		clientReq.Header.Add(key, value)
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

	err = json.Unmarshal(body, r.Response)
	return resp, err
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
