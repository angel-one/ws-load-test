package middlewares

import (
	"github.com/angel-one/go-utils/constants"
	"github.com/angel-one/go-utils/log"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"time"
)

// LoggerMiddlewareOptions is the set of configurable allowed for log
type LoggerMiddlewareOptions struct {
	NotLogQueryParams  bool
	NotLogHeaderParams bool
}

// Logger is the middleware to be used for logging the request and response information
// This should be the first middleware to be added, in case the recovery middleware is not being used.
// Otherwise, it should be the second one, just after the recovery middleware.
func Logger(options LoggerMiddlewareOptions) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.GetHeader(constants.RequestIDHeader)

		if id == "" {
			// get a unique id
			uid, err := uuid.NewUUID()
			if err == nil {
				id = uid.String()
			}
		}

		// apply the id in the context
		ctx.Set(constants.IDLogParam, id)

		// Start timer
		start := time.Now()
		path := ctx.Request.URL.Path

		// log the initial request
		event := log.Info(ctx).
			Str(constants.IDLogParam, id).
			Time(constants.StartTimeLogParam, start).
			Str(constants.MethodLogParam, ctx.Request.Method).
			Str(constants.PathLogParam, path)
		if !options.NotLogQueryParams {
			event.Str(constants.QueryLogParam, ctx.Request.URL.RawQuery)
		}
		if !options.NotLogHeaderParams {
			event.Interface(constants.HeaderLogParams, ctx.Request.Header)
		}
		event.Send()

		// Process request
		ctx.Next()

		// stop timer
		end := time.Now()
		latency := end.Sub(start)

		// log the final response details
		log.Info(ctx).
			Str(constants.IDLogParam, id).
			Int(constants.StatusCodeLogParam, ctx.Writer.Status()).
			Time(constants.EndTimeLogParam, end).
			Dur(constants.LatencyLogParam, latency).
			Str(constants.ClientIPLogParam, ctx.ClientIP()).
			Str(constants.ErrorLogParam, ctx.Errors.ByType(gin.ErrorTypePrivate).String()).
			Send()
	}
}
