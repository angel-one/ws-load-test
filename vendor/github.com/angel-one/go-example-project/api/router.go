package api

import (
	"github.com/angel-one/go-example-project/constants"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

// GetRouter is used to get the router configured with the middlewares and the routes
func GetRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	router := gin.New()
	router.Use(middlewares...)
	router.Use(gin.Recovery())

	// configure swagger
	router.GET(constants.SwaggerRoute, ginSwagger.WrapHandler(swaggerFiles.Handler))

	// configure actuator
	router.GET(constants.ActuatorRoute, actuator)

	// adding api
	router.POST(constants.FullNameRoute, fullName)
	router.GET(constants.MoxyRoute, moxy)
	router.POST(constants.CreateCounterRoute, createCounter)
	router.PUT(constants.IncrementCounterRoute, incrementCounter)
	router.POST(constants.DecrementCounterRoute, decrementCounter)
	router.GET(constants.CurrentCountRoute, currentCount)


	// async jobs api
	router.POST(constants.MathsSubmitJobRoute, submitMathsJob)
	router.GET(constants.MathsGetJobStatusRoute, getJobStatus)

	return router
}
