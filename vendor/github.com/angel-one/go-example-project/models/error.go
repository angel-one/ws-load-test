package models

// ErrorResponse is the common error response body to be used in case of any error
type ErrorResponse struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}
