package errors

// Http error is a object error generic to excepetion http requests
type Http struct {
	StatusCode int      `json:"status_code"`
	Message    string   `json:"message"`
	Errors     []string `json:"errors"`
}

// NewHttpError is a factory function that generate a new Http error
func NewHttpError(statusCode int, message string, errors []string) *Http {
	return &Http{
		StatusCode: statusCode,
		Message:    message,
		Errors:     errors,
	}
}
