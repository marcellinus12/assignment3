package routers

import (
	"assignment3/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/status", controllers.UpdateStatus)

	return router
}
