package wrecker

type ResponseError struct {
	StatusCode int
	StatusText string
}

func (r ResponseError) Error() string {
	return r.StatusText
}
