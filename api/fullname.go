package api

import (
	"github.com/angel-one/go-example-project/business"
	"github.com/angel-one/go-example-project/constants"
	"github.com/angel-one/go-example-project/models"
	"github.com/angel-one/go-utils/log"
	"github.com/gin-gonic/gin"
	"net/http"
)

// fullName godoc
// @Summary Gets the full name from the first name and last name
// @Description Gets the full name from the first name and last name
// @ID fullName
// @Tags fullName
// @Accept  json
// @Produce  json
// @Param request body models.FullNameRequest true "first name and last name"
// @Success 200 {object} models.FullNameResponse
// @Failure 400 {object} models.ErrorResponse
// @Router /fullName [post]
func fullName(ctx *gin.Context) {
	var request models.FullNameRequest

	// here using ctx.BindJSON(), we will lose the control of sending error response according to our structure
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		// send error response
		// this is a scenario of bad request body
		log.Error(ctx).Stack().Err(err).Msg("error binding request body")
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:        constants.RequestBodyBindError,
			Description: err.Error(),
		})
		return
	}

	// now validate the request body
	err = request.Validate()
	if err != nil {
		// send error response
		// this is again a scenario of bad request body
		log.Error(ctx).Stack().Err(err).Msg("error validating request body")
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			Code:        constants.RequestBodyValidationError,
			Description: err.Error(),
		})
		return
	}

	// now we have the request body
	// apply our business logic to this request
	response := business.GetFullName(request)

	// now send the response
	ctx.JSON(http.StatusOK, response)
}
