package routes

import (
	"github.com/danteay/lanago/api/handlers"
	"github.com/danteay/lanago/api/config"
	"github.com/gin-gonic/gin"
)

// InitRoutes register all the endpoits for the application.
// It needs a gin.Engine pointer to bind the routes and config.ServiceClients
// as context of the application and its dependencies
func InitRoutes(app *gin.Engine, context *config.ServiceClients) {
	app.GET("/ping", handlers.PingHandler())

	app.POST("/basket", handlers.NewBasket(context))

	app.DELETE("/basket/:basketId", handlers.DeleteBasket(context))
	app.POST("/basket/:basketId", handlers.AddBasketItem(context))
	app.GET("/basket/:basketId", handlers.GetBasket(context))
}
