package routes

import (
	"lanago/api/handlers"
	"lanago/config"

	"github.com/gin-gonic/gin"
)

// InitRoutes register all the endpoits for the application.
// It needs point of gin.Engine to bind the routes
func InitRoutes(app *gin.Engine, context *config.ServiceClients) {
	app.GET("/ping", handlers.PingHandler())

	app.POST("/basket", handlers.NewBasket(context))

	app.DELETE("/basket/:basketId", handlers.DeleteBasket(context))
	app.POST("/basket/:basketId", handlers.AddBasketItem(context))
	app.GET("/basket/:basketId", handlers.GetBasket(context))
}
