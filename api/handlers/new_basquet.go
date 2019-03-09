package handlers

import (
	"encoding/json"
	"lanago/api/models"
	"lanago/config"
	"time"

	rest "github.com/danteay/ginrest"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// NewBasquet is the action handler the registration of a new basquet.
// It generates a new UUID v4 for the ID of the basquent and storage on Redis,
// Then return an object response with the created data
func NewBasket(context *config.ServiceClients) func(c *gin.Context) {
	return func(c *gin.Context) {
		u := c.Request.RequestURI
		r := rest.New(u, "").SetGin(c)

		basket := models.Basket{}
		basket.ID = uuid.Must(uuid.NewV4()).String()
		basket.Items = []models.Product{}

		go func() {
			j, err := json.Marshal(&basket)

			if err != nil {
				context.Logger.Error("lanago.post.NewBasket", err.Error())
				return
			}

			err = context.Redis.Set(basket.ID, string(j), time.Hour*24).Err()

			if err != nil {
				context.Logger.Error("lanago.post.NewBasket", err.Error())
			}
		}()

		r.Res(200, rest.Payload{
			"status": "success",
			"object": "lanago.post.NewBasquet",
			"data":   basket,
		}, "basket created")
		return
	}
}
