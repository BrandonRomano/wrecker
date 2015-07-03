package interceptors

import (
	"github.com/brandonromano/wrecker"
)

// Authorization is a sample Interceptor that adds a HTTP "Authorization" header to every request.
func Authorization(auth string) wrecker.Interceptor {

	return wrecker.Interceptor{

		// This is executed on every Request before its sent to the server
		Request: func(r *wrecker.Request) error {

			r.Header("Authorization", auth)
			return nil
		},
	}
}
