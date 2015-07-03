package interceptors

import (
	"fmt"
	"github.com/brandonromano/wrecker"
	"net/http"
	"strconv"
)

// Debug is a sample Interceptor that adds debugging output to every request.
func Debug() wrecker.Interceptor {

	return wrecker.Interceptor{

		// This is executed on every Request before its sent to the server
		Request: func(r *wrecker.Request) error {

			return nil
		},

		RawRequest: func(r *http.Request) error {

			fmt.Println("")
			fmt.Println("HTTP Request")
			fmt.Println("-------------")
			fmt.Println("URL: ", r.URL.String())
			fmt.Println("Method: ", r.Method)
			return nil
		},

		RawResponse: func(r *http.Response, body []byte) error {

			fmt.Println("")
			fmt.Println("HTTP Response")
			fmt.Println("-------------")

			fmt.Println("Status Code: ", strconv.Itoa(r.StatusCode))
			fmt.Println("Status: ", r.Status)
			fmt.Println("Headers:")

			for i := range r.Header {
				fmt.Println("- ", i, ": ", r.Header.Get(i))
			}

			fmt.Println("Body: ", string(body))
			fmt.Println("")

			return nil
		},
	}
}
