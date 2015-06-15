package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Success    bool        `json:"success"`
	StatusCode int         `json:"status_code"`
	StatusText string      `json:"status_text"`
	Content    interface{} `json:"content"`
}

func (response *Response) Init() *Response {
	response.StatusCode = http.StatusOK
	return response
}

func (response *Response) Output(writer http.ResponseWriter) {
	// Determining Unknowns
	response.Success = response.StatusCode == http.StatusOK
	response.StatusText = http.StatusText(response.StatusCode)

	// Writing to the ResponseWriter
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	writer.WriteHeader(response.StatusCode)
	fmt.Fprint(writer, toJSON(response))
}

func toJSON(i interface{}) string {
	json, err := json.Marshal(i)
	if err != nil {
		return "Error converting response to JSON"
	}
	return string(json)
}
