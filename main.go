package main

import (
	"github.com/gin-gonic/gin"
	"github.com/samuelIkoli/fun_numbers/routes"
)

func main(){


	server:= gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")

}