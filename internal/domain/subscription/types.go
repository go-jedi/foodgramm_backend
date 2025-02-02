package subscription

// ErrorResponse represents a generic error response.
type ErrorResponse struct {
	Error  string `json:"error"`
	Detail string `json:"detail"`
}
