package api

import (
	"errors"
	"github.com/angel-one/go-example-project/business"
	"github.com/angel-one/go-example-project/constants"
	"github.com/angel-one/go-example-project/models"
	"github.com/angel-one/go-utils/log"
	"github.com/gin-gonic/gin"
	"net/http"
)

// createCounter godoc
// @Summary Creates a new counter
// @Description Creates a new counter
// @ID createCounter
// @Tags counter
// @Produce  json
// @Param key query string true "counter key"
// @Success 201
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /counter/create [post]
func createCounter(ctx *gin.Context) {
	// get the key and validate
	key := ctx.Query(constants.CounterKey)
	err := validateCounterKey(key)
	if err != nil {
		sendCounterRequestValidationError(ctx, err)
		return
	}

	// now create a new counter
	err = business.CreateCounter(key)
	if err != nil {
		sendCounterInternalServerError(ctx, err)
		return
	}

	// send success response
	ctx.Status(http.StatusCreated)
}

// incrementCounter godoc
// @Summary Increment an existing counter
// @Description Increment an existing counter
// @ID incrementCounter
// @Tags counter
// @Produce  json
// @Param key query string true "counter key"
// @Success 200
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /counter/increment [put]
func incrementCounter(ctx *gin.Context) {
	// get the key and validate
	key := ctx.Query(constants.CounterKey)
	err := validateCounterKey(key)
	if err != nil {
		sendCounterRequestValidationError(ctx, err)
		return
	}

	// now increment the counter
	err = business.IncrementCounter(key)
	if err != nil {
		sendCounterInternalServerError(ctx, err)
	}
}

// decrementCounter godoc
// @Summary Decrement an existing counter
// @Description Decrement an existing counter
// @ID decrementCounter
// @Tags counter
// @Produce  json
// @Param key query string true "counter key"
// @Success 200
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /counter/decrement [put]
func decrementCounter(ctx *gin.Context) {
	// get the key and validate
	key := ctx.Query(constants.CounterKey)
	err := validateCounterKey(key)
	if err != nil {
		sendCounterRequestValidationError(ctx, err)
		return
	}

	// now decrement the counter
	err = business.DecrementCounter(ctx, key)
	if err != nil {
		sendCounterInternalServerError(ctx, err)
	}
}

// currentCount godoc
// @Summary Get the current value of counter
// @Description Get the current value of counter
// @ID currentCount
// @Tags counter
// @Produce  json
// @Param key query string true "counter key"
// @Success 200 {object} models.CounterResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /counter/current [get]
func currentCount(ctx *gin.Context) {
	// get the key and validate
	key := ctx.Query(constants.CounterKey)
	err := validateCounterKey(key)
	if err != nil {
		sendCounterRequestValidationError(ctx, err)
		return
	}

	count, err := business.CurrentCount(key)
	if err != nil {
		sendCounterInternalServerError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, models.CounterResponse{
		Key:   key,
		Count: count,
	})
}

func validateCounterKey(key string) error {
	if key == "" {
		return errors.New("invalid key provided, cannot be empty")
	}
	return nil
}

func sendCounterRequestValidationError(ctx *gin.Context, err error) {
	log.Error(ctx).Stack().Err(err).Msg("invalid counter key")
	ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
		Code:        constants.RequestValidationError,
		Description: err.Error(),
	})
}

func sendCounterInternalServerError(ctx *gin.Context, err error) {
	log.Error(ctx).Stack().Err(err).Msg("unable to work with counter")
	ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
		Code:        constants.DatabaseFailureError,
		Description: err.Error(),
	})
}
