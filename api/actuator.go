package api

import (
	"github.com/angel-one/ws-load-test/constants"
	"github.com/gin-gonic/gin"
	goActuator "github.com/sinhashubham95/go-actuator"
)

var (
	actuatorHandler = goActuator.GetActuatorHandler(&goActuator.Config{
		Name:    constants.ApplicationName,
		Version: "",
	})
)

func actuator(ctx *gin.Context) {
	actuatorHandler(ctx.Writer, ctx.Request)
}
