package api

import (
	"github.com/gin-gonic/gin"
)

func MetricRoutes(group *gin.RouterGroup) {
	group.GET("", welcomeController)
	group.GET("latency", latencyController)
	group.GET("send", sendController)
	group.GET("receive", receiveController)
	group.GET("error", errorController)
	group.GET("connection", connectionController)
	group.GET("dashboard", dashboardController)

}


