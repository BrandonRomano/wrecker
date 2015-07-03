package interceptors

import (
	"fmt"
	"github.com/brandonromano/wrecker"
	"net/http"
)

// Debug is a sample Interceptor that adds debugging output to every request.
func Debug(auth string) wrecker.Interceptor {

	return wrecker.Interceptor{

		// This is executed on every Request before its sent to the server
		WreckerRequest: func(r *wrecker.Request) error {

			return nil
		},

		RawRequest: func(r *http.Request) error {

			return nil
		},

		RawResponse: func(r *http.Response) error {

			return nil
		},
	}
}
