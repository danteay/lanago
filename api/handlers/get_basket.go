package handlers

import (
	"encoding/json"
	"errors"

	"lanago/api/models"
	"lanago/api/services/discounts"
	"lanago/config"

	rest "github.com/danteay/ginrest"
	"github.com/gin-gonic/gin"
)

func GetBasket(context *config.ServiceClients) func(c *gin.Context) {
	return func(c *gin.Context) {
		u := c.Request.RequestURI
		r := rest.New(u, "").SetGin(c)

		id := c.Param("basketId")

		j, err := context.Redis.Get(id).Result()

		if err != nil && err.Error() == "redis: nil" {

			context.Logger.Error("lanago.get.GetBasket", err.Error())
			getBasketErrorResponse(r, errors.New("basket not foud"), 404)
			return
		} else if err != nil {
			context.Logger.Error("lanago.get.GetBasket", err.Error())
			getBasketErrorResponse(r, err, 500)
			return
		}

		var basket models.Basket

		err = json.Unmarshal([]byte(j), &basket)

		if err != nil {
			context.Logger.Error("lanago.get.GetBasket", err.Error())
			getBasketErrorResponse(r, err, 500)
			return
		}

		setTotalBasket(&basket)

		r.Res(200, rest.Payload{
			"status": "success",
			"object": "lanago.get.GetBasket",
			"data":   basket,
		}, "current basket")
	}
}

func getBasketErrorResponse(r *rest.IO, err error, code int) {
	r.Res(code, rest.Payload{
		"status": "error",
		"object": "lanago.get.GetBasket",
	}, err.Error())
}

func setTotalBasket(b *models.Basket) {
	b.Total = 0

	for _, i := range b.Items {
		b.Total = b.Total + (i.Price * float64(i.Qty))
	}

	applyDiscounts(b)
}

// applyDiscounts is the wrapper funtion to apply all the available discounts
// at the moment for the current basket
func applyDiscounts(b *models.Basket) {
	discounts.TwoXOne(b, func(elem interface{}) bool {
		i := elem.(models.Product)
		return i.Code == "VOUCHER" && i.Qty >= 2
	})

	discounts.BulkPurchase(b, 19, func(elem interface{}) bool {
		i := elem.(models.Product)
		return i.Code == "TSHIRT" && i.Qty >= 3
	})
}