package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/samuelIkoli/fun_numbers/controllers"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Deployed and Running fun nummbers app"})
	})

	router.GET("/task", controller.Task)
	router.GET("/test", controller.Test)
	router.GET("/api", controller.Test)
	router.GET("/api/classify-number", controller.Numerate)
	router.GET("/ping", controller.Ping)
}