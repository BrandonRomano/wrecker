package wrecker

import (
	"errors"
	"net/http"
	"net/url"
)

type Request struct {
	HttpVerb      string
	Endpoint      string
	Response      interface{}
	Params        url.Values
	Body          interface{}
	Headers       map[string]string
	WreckerClient *Wrecker
}

func (r *Request) WithParam(key, value string) *Request {
	r.Params.Add(key, value)
	return r
}

func (r *Request) WithHeader(key, value string) *Request {
	r.Headers[key] = value
	return r
}

func (r *Request) WithBody(body interface{}) *Request {
	r.Body = body
	return r
}

func (r *Request) Into(response interface{}) *Request {
	r.Response = response
	return r
}

func (r *Request) Execute() (*http.Response, error) {
	switch r.HttpVerb {

	case GET, POST, PUT, DELETE:
		return r.WreckerClient.sendRequest(r)

	default:
		return nil, errors.New("Must use a valid HTTP verb")
	}
}

func (r *Request) URL() string {
	result := r.WreckerClient.BaseURL + r.Endpoint

	if (r.HttpVerb == "GET") && (len(r.Params) > 0) {
		result += "?" + r.Params.Encode()
	}
	return result
}
