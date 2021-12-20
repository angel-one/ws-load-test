package business

import "github.com/angel-one/go-example-project/models"

// GetFullName is the business logic for deciding the full name based on the request
func GetFullName(request models.FullNameRequest) models.FullNameResponse {
	fullName := request.FirstName
	if request.LastName != "" {
		fullName = fullName + " " + request.LastName
	}
	return models.FullNameResponse{
		FullName: fullName,
	}
}
