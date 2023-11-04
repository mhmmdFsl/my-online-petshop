package http_error

type HttpError struct {
	Message    string
	StatusCode int
	Code       string
}

func (h HttpError) Error() string {
	return h.Message
}

func NewFailed(m string, s int) HttpError {
	return HttpError{
		Message:    m,
		StatusCode: s,
		Code:       "FAILED",
	}
}

func NewError(m string) HttpError {
	return HttpError{
		Message:    m,
		StatusCode: 500,
		Code:       "FAILED",
	}
}
