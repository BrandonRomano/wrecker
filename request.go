package wrecker

import (
	"errors"
	"net/url"
)

type WreckerRequest struct {
	HttpVerb      string
	Endpoint      string
	Response      interface{}
	Params        url.Values
	Body          interface{}
	Headers       map[string]string
	WreckerClient *Wrecker
}

func (r *WreckerRequest) WithParam(key, value string) *WreckerRequest {
	r.Params.Add(key, value)
	return r
}

func (r *WreckerRequest) WithHeader(key, value string) *WreckerRequest {
	r.Headers[key] = value
	return r
}

func (r *WreckerRequest) WithBody(body interface{}) *WreckerRequest {
	r.Body = body
	return r
}

func (r *WreckerRequest) Into(response interface{}) *WreckerRequest {
	r.Response = response
	return r
}

func (r *WreckerRequest) Execute() error {

	switch r.HttpVerb {

	case GET, POST, PUT, DELETE:
		return r.WreckerClient.sendRequest(r)

	default:
		return errors.New("Must use a valid HTTP verb")
	}
}

func (r *WreckerRequest) URL() string {

	result := r.WreckerClient.BaseURL + r.Endpoint

	if (r.HttpVerb == "GET") && (len(r.Params) > 0) {
		result += "?" + r.Params.Encode()
	}

	return result
}
