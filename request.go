package wrecker

import (
	"errors"
	"net/http"
	"net/url"
)

type BasicAuthInfo struct {
	Username string
	Password string
}

type Request struct {
	HttpVerb      string
	Endpoint      string
	Response      interface{}
	URLParams     url.Values
	FormParams    url.Values
	HttpBody      interface{}
	Headers       map[string]string
	Cookies       []*http.Cookie
	BasicAuthInfo *BasicAuthInfo
	WreckerClient *Wrecker
}

func (r *Request) Cookie(cookie *http.Cookie) *Request {
	r.Cookies = append(r.Cookies, cookie)
	return r
}

func (r *Request) Header(key, value string) *Request {
	r.Headers[key] = value
	return r
}

func (r *Request) URLParam(key, value string) *Request {
	r.URLParams.Add(key, value)
	return r
}

func (r *Request) FormParam(key, value string) *Request {
	r.FormParams.Add(key, value)
	return r
}

func (r *Request) Body(body interface{}) *Request {
	r.HttpBody = body
	return r
}

func (r *Request) SetBasicAuth(username, password string) *Request {
	r.BasicAuthInfo = &BasicAuthInfo{
		Username: username,
		Password: password,
	}
	return r
}

func (r *Request) Into(response interface{}) *Request {
	r.Response = response
	return r
}

func (r *Request) Execute() (*http.Response, error) {
	// Calling interceptor if we have one
	if r.WreckerClient.RequestInterceptor != nil {
		r.WreckerClient.RequestInterceptor(r)
	}

	// Sending Request
	switch r.HttpVerb {

	case GET, POST, PUT, DELETE, PATCH:
		return r.WreckerClient.sendRequest(r)

	default:
		return nil, errors.New("Must use a valid HTTP verb")
	}
}

func (r *Request) URL() string {
	result := r.WreckerClient.BaseURL + r.Endpoint

	if len(r.URLParams) > 0 {
		result += "?" + r.URLParams.Encode()
	}

	return result
}
