package api

import (
	"github.com/angel-one/ws-load-test/constants"
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

	router.GET(constants.SwaggerRoute, ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET(constants.ActuatorRoute, actuator)

	reportRouter := router.Group("metric")
	MetricRoutes(reportRouter)

	return router
}
