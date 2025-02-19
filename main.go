package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
	"github.com/samuelIkoli/fun_numbers/routes"
)

func main(){

	c := cron.New()
	c.AddFunc("*/550 * * * *", func() {
		fmt.Println("Cronning")
		// http.Get("https://fun-numbers.onrender.com/")

	})
	c.Start()

	server:= gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")

}