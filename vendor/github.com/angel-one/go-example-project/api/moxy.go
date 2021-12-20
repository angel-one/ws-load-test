package api

import (
	"github.com/angel-one/go-example-project/business"
	"github.com/angel-one/go-example-project/constants"
	"github.com/angel-one/go-example-project/models"
	"github.com/angel-one/go-utils/log"
	"github.com/gin-gonic/gin"
	"net/http"
)

// moxy godoc
// @Summary Get the moxy response
// @Description Get the moxy response
// @ID moxy
// @Tags moxy
// @Produce  json
// @Success 200 {object} models.MoxyResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /moxy [get]
func moxy(ctx *gin.Context) {
	data, err := business.GetMoxy(ctx)
	if err != nil {
		log.Error(ctx).Stack().Err(err).Msg("error getting moxy")
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Code:        constants.ExternalServiceFailureError,
			Description: err.Error(),
		})
		return
	}

	// now handle success scenario
	ctx.JSON(http.StatusOK, data)
}
