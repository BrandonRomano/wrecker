package wrecker

import (
	"encoding/json"
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

func (w *Wrecker) sendRequest(verb string, requestURL string, headers map[string]string, bodyParams url.Values, response interface{}) (*http.Response, error) {
	// Preparing Request
	paramsReader := strings.NewReader(bodyParams.Encode())
	clientReq, err := http.NewRequest(verb, requestURL, paramsReader)
	if err != nil {
		return nil, err
	}
	clientReq.Header.Add("Content-Type", w.DefaultContentType)

	// Add headers to clientReq
	for index, value := range headers {
		clientReq.Header.Add(index, value)
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
	err = json.Unmarshal(body, &response)
	return resp, err
}
