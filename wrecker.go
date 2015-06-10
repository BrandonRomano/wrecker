package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Wrecker struct {
	BaseURL    string
	HttpClient *http.Client
}

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
)

func (w *Wrecker) Get(endpoint string, params url.Values, response interface{}) error {
	return w.SendRequest(GET, endpoint, params, response)
}

func (w *Wrecker) Post(endpoint string, params url.Values, response interface{}) error {
	return w.SendRequest(POST, endpoint, params, response)
}

func (w *Wrecker) Put(endpoint string, params url.Values, response interface{}) error {
	return w.SendRequest(PUT, endpoint, params, response)
}

func (w *Wrecker) Delete(endpoint string, response interface{}) error {
	return w.SendRequest(DELETE, endpoint, nil, response)
}

func (w *Wrecker) SendRequest(verb string, endpoint string, params url.Values, response interface{}) error {
	// Preparing Request
	requestURL := strings.Join([]string{w.BaseURL, endpoint}, "")
	paramsReader := strings.NewReader(params.Encode())
	clientReq, err := http.NewRequest(verb, requestURL, paramsReader)
	if err != nil {
		return err
	}
	clientReq.Header.Add("Content-Type", "application/x-www-form-urlencoded")

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
	err = json.Unmarshal(body, &response)
	return err
}
