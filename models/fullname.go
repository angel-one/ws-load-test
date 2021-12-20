package models

import "errors"

// FullNameRequest is the request body for the full name api
type FullNameRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// FullNameResponse is the response body for the full name api
type FullNameResponse struct {
	FullName string `json:"name"`
}

// Validate is used to validate the request body
func (r FullNameRequest) Validate() error {
	if r.FirstName == "" {
		return errors.New("first name cannot be empty")
	}
	return nil
}
