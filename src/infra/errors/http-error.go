package errors

// Http error is a struct error generic to excepetion http requests
type Http struct {
	StatusCode int      `json:"status_code"`
	Message    string   `json:"message"`
	Errors     []string `json:"errors"`
}
