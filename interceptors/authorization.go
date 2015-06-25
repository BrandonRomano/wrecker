package interceptors

import (
	"github.com/BrandonRomano/wrecker"
)

// Authorization generates a wrecker.RequestInterceptorFunc that adds
// the "Authorization" header into every HTTP request.
func Authorization(auth string) wrecker.RequestInterceptorFunc {

	return func(r *wrecker.WreckerRequest) error {

		r.Header("Authorization", auth)

		return nil
	}
}
