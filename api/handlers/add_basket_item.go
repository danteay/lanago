package handlers

import (
	"encoding/json"
	"time"

	"lanago/api/models"
	"lanago/config"

	rest "github.com/danteay/ginrest"
	"github.com/gin-gonic/gin"
)

// addBasketItemRequest is the request data pf the endpoint
type addBasketItemRequest struct {
	Code string `json:"code,required"`
}

// AddBasketItem is the action handler to add a new Product to an specific
// basket. If the basket is not found it fails, and if is finded, will add the
// product is asyc way.
func AddBasketItem(context *config.ServiceClients) func(c *gin.Context) {
	return func(c *gin.Context) {
		u := c.Request.RequestURI
		r := rest.New(u, "").SetGin(c)

		id := c.Param("basketId")

		data, err := context.Redis.Get(id).Result()

		if err != nil {
			context.Logger.Error("lanago.post.AddBasketItem", err.Error())
			addBasketItemErrorResponse(r, err, 500)
			return
		}

		var req addBasketItemRequest
		err = c.BindJSON(&req)

		if err != nil {
			context.Logger.Error("lanago.post.AddBasketItem", err.Error())
			addBasketItemErrorResponse(r, err, 500)
			return
		}

		var basket models.Basket

		err = json.Unmarshal([]byte(data), &basket)

		if err != nil {
			context.Logger.Error(err.Error())
			return
		}

		finded := false

		for i, item := range basket.Items {
			if item.Code == req.Code {
				item.Qty++
				basket.Items[i] = item
				finded = true
				break
			}
		}

		var aux *models.Product

		if !finded {
			aux, err = context.ProductsService.FindProduct(req.Code)

			if err != nil {
				context.Logger.Error(err.Error())
				addBasketItemErrorResponse(r, err, 422)
				return
			}

			aux.Qty = 1

			basket.Items = append(basket.Items, *aux)
		}

		go saveBasket(&basket, context)

		r.Res(202, rest.Payload{
			"status": "success",
			"object": "lanago.post.AddBasketItem",
		}, "action performed")
		return
	}
}

// addBasketItemErrorResponse is a wrapper for a general error response inside
// the main handler AddBasketItem
func addBasketItemErrorResponse(r *rest.IO, err error, code int) {
	r.Res(code, rest.Payload{
		"status": "error",
		"object": "lanago.post.AddBasketItem",
	}, err.Error())
}

// processBasket is the asyc task that will add the new product to the current
// basket. This find for a previous insertion of the same product into the
// basket, if is not found, then creates a new record, update the current total
// and finally apply the available discounts an storage the result basket on
// redis again
func processBasket(b *models.Basket, c *config.ServiceClients) {
	// applyDiscounts(b)
	saveBasket(b, c)
}

// SaveBasket storage the computed information into Redis
func saveBasket(b *models.Basket, c *config.ServiceClients) {
	j, err := json.Marshal(b)

	if err != nil {
		c.Logger.Error(err.Error())
		return
	}

	err = c.Redis.Set(b.ID, string(j), time.Hour).Err()

	if err != nil {
		c.Logger.Error(err.Error())
		return
	}
}
