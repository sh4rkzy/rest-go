package health

import (
	"github.com/gin-gonic/gin"	
	"github.com/sh4rkzy/rest-go/src/modules/health/controllers"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/health", controllers.HealthCheck)	
}