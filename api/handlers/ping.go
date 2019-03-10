package handlers

import (
	rest "github.com/danteay/ginrest"
	"github.com/gin-gonic/gin"
)

// PingHandler is a simple get endpoint that can be used for healt check
func PingHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		u := c.Request.RequestURI
		r := rest.New(u, "").SetGin(c)

		r.Res(200, rest.Payload{
			"status": "success",
			"object": "lanago.get.ping",
			"code":   200,
		}, "pong...")
		return
	}
}
