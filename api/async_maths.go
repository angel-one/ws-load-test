package api

import (
	"github.com/angel-one/go-example-project/business"
	"github.com/angel-one/go-example-project/constants"
	"github.com/angel-one/go-example-project/models"
	"github.com/angel-one/go-utils/log"
	"github.com/gin-gonic/gin"
	"net/http"
)

// submitMathsJob godoc
// @Summary Does a maths op in a async way
// @Description Returns a op b
// @ID asyncMaths
// @Tags asyncMaths
// @Accept  json
// @Produce  json
// @Param request body models.MathsRequest true "a op b"
// @Success 200 {object} models.MathsResponse
// @Failure 400 {object} models.MathsResponse
// @Router /maths [post]
func submitMathsJob(ctx *gin.Context) {
	var request models.MathsRequest

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

	// TODO  validate the request body

	// now do the app
	response := business.SubmitMathsJob(&request)

	// now send the response
	// TODO validation
	ctx.JSON(http.StatusOK, response)
}


// getJobStatus godoc
// @Summary Get the status (and result if completed) of the job
// @Description GGet the status (and result if completed) of the job
// @ID getJobStatus
// @Tags asyncMaths
// @Produce  json
// @Param key query string true "counter key"
// @Success 200 {object} jobs.JobResult
// @Failure 400 {object} jobs.JobResult
// @Failure 500 {object} jobs.JobResult
// @Router /maths/{id} [get]
func getJobStatus(ctx *gin.Context) {
	// get the key and validate
	key := ctx.Param("jobid")

	res, err := business.GetMathJobStatus(key)
	if err != nil {
		log.Error(ctx).Stack().Err(err).Msg("unable to get job status ")
		ctx.JSON(http.StatusNotFound, res)
		return
	}

	ctx.JSON(http.StatusOK, res)
}