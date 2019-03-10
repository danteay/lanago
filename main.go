package main

import (
	"log"
	"os"

	"github.com/danteay/lanago/api/routes"
	"github.com/danteay/lanago/config"

	"github.com/gin-gonic/gin"
)

var port string

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Panic error: ", r)
		}
	}()

	context := new(config.ServiceClients)
	context.Init()

	app := gin.New()
	app.Use(gin.Logger())
	app.Use(gin.Recovery())

	routes.InitRoutes(app, context)

	if port = os.Getenv("PORT"); port == "" {
		port = "8080"
	}

	app.Run(":" + port)
}
