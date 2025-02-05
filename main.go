package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
	"github.com/samuelIkoli/fun_numbers/routes"
)

func main(){

	c := cron.New()
	c.AddFunc("*/500 * * * *", func() {
		fmt.Println("Cronning")
		http.Get("https://hng12-be-0.onrender.com/")

	})
	c.Start()

	server:= gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")

}