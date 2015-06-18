package wrecker

import (
	"errors"
	"net/url"
	"strings"
)

type WreckerRequest struct {
	HttpVerb      string
	Endpoint      string
	Response      interface{}
	Params        url.Values
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

func (r *WreckerRequest) Into(response interface{}) *WreckerRequest {
	r.Response = response
	return r
}

func (r *WreckerRequest) Execute() error {
	switch r.HttpVerb {
	case GET:
		return r.executeGet()
	case POST:
		return r.executePost()
	case PUT:
		return r.executePut()
	case DELETE:
		return r.executeDelete()
	default:
		return errors.New("Must use a valid HTTP verb")
	}
}

func (r *WreckerRequest) executeGet() error {
	requestURL := strings.Join([]string{r.WreckerClient.BaseURL, r.Endpoint, "?", r.Params.Encode()}, "")
	return r.WreckerClient.sendRequest(GET, requestURL, r.Headers, nil, r.Response)
}

func (r *WreckerRequest) executePost() error {
	requestURL := strings.Join([]string{r.WreckerClient.BaseURL, r.Endpoint}, "")
	return r.WreckerClient.sendRequest(POST, requestURL, r.Headers, r.Params, r.Response)
}

func (r *WreckerRequest) executePut() error {
	requestURL := strings.Join([]string{r.WreckerClient.BaseURL, r.Endpoint}, "")
	return r.WreckerClient.sendRequest(PUT, requestURL, r.Headers, r.Params, r.Response)
}

func (r *WreckerRequest) executeDelete() error {
	requestURL := strings.Join([]string{r.WreckerClient.BaseURL, r.Endpoint}, "")
	return r.WreckerClient.sendRequest(DELETE, requestURL, r.Headers, nil, r.Response)
}
