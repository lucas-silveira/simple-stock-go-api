package errors

// Http error is a struct error generic to excepetion http requests
type Http struct {
	StatusCode int
	Message    string
}
