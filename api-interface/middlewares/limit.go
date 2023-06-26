package middlewares

import (
	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LimitSource() gin.HandlerFunc {
	return func(c *gin.Context) {

		if tag := c.Request.Header.Get("authorization"); tag == "open-api-gateway" {
			c.Next()
			return
		}

		c.JSON(http.StatusForbidden, nil)
		c.Abort()
	}
}