package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/samuelIkoli/fun_numbers/handlers"
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
	router.GET("/integration.json", controller.Webhook)
	router.GET("/integration", controller.Webhook2)
	router.GET("/tick", controller.Get_symbols)
	router.POST("/tick", controller.Get_symbols)
}