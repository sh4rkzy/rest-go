package main 

import (
	"github.com/sh4rkzy/rest-go/src/modules/health/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router  := gin.Default()
	router.Group("api/v1")
	health.RegisterRoutes(router)
	router.Run(":8080")
}