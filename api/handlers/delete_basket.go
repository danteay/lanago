package handlers

import (
	"github.com/danteay/lanago/config"

	rest "github.com/danteay/ginrest"
	"github.com/gin-gonic/gin"
)

// DeleteBasket is the action that handle the detletion of an existing basket.
// It receives a basket ID from the URL amd search on Redis to delete it on
// asyc way
func DeleteBasket(context *config.ServiceClients) func(c *gin.Context) {
	return func(c *gin.Context) {
		u := c.Request.RequestURI
		r := rest.New(u, "").SetGin(c)

		id := c.Param("basketId")

		go func() {
			err := context.Redis.Del(id).Err()
			if err != nil {
				context.Logger.Error("lana.delete.DeleteBasket", err.Error())
			}
		}()

		r.Res(202, rest.Payload{
			"status": "success",
			"object": "lagago.delete.DeleteBasket",
		}, "action performed")
		return
	}
}
