package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/robfig/cron"
	"github.com/samuelIkoli/fun_numbers/routes"
)

func main(){
	err := godotenv.Load()
	if err != nil {
		// log.Fatal("Error loading .env file")
	}
	c := cron.New()
	c.AddFunc("*/550 * * * *", func() {
		fmt.Println("Cronning")
		http.Get("https://fun-numbers.onrender.com/")

	})
	c.Start()

	server:= gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")

}