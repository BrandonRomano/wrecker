package interceptors

import (
	"fmt"
	"github.com/brandonromano/wrecker"
	// "io/ioutil"
	"net/http"
	"strconv"
)

// Debug is a sample Interceptor that adds debugging output to every request.
func Debug() wrecker.Interceptor {

	return wrecker.Interceptor{

		// This is executed on every Request before its sent to the server
		WreckerRequest: func(r *wrecker.Request) error {

			fmt.Println("")
			fmt.Println("Wrecker Request")
			fmt.Println("-------------")
			fmt.Println("Method: ", r.HttpVerb)
			fmt.Println("URL: ", r.URL())

			fmt.Println("Headers:")

			for i := range r.Headers {
				fmt.Println("- ", i, ": ", r.Headers[i])
			}

			fmt.Println("URLParams:")

			for i := range r.URLParams {
				fmt.Println("- ", i, ": ", r.URLParams[i])
			}

			fmt.Println("FormParams:")

			for i := range r.FormParams {
				fmt.Println("- ", i, ": ", r.FormParams[i])
			}

			fmt.Println("")

			return nil
		},

		HTTPRequest: func(r *http.Request) error {

			fmt.Println("")
			fmt.Println("HTTP Request")
			fmt.Println("-------------")
			fmt.Println("Method: ", r.Method)
			fmt.Println("URL: ", r.URL.String())
			// fmt.Println("Body: ", string(ioutil.ReadAll(r.Body)))
			// fmt.Println("Body (Deux): ", string(ioutil.ReadAll(r.Body)))
			// fmt.Println("Body (Trois): ", string(ioutil.ReadAll(r.Body)))
			fmt.Println("")

			return nil
		},

		HTTPResponse: func(r *http.Response, body []byte) error {

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
