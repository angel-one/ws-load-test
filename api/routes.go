package api

import (
	"github.com/gin-gonic/gin"
)

func MetricRoutes(group *gin.RouterGroup) {
	group.GET("/welcome", welcomeController)
	group.GET("/latency", latencyController)
}


