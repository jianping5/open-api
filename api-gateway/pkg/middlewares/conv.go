package middlewares

import (
	"fmt"
	// "net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func Conv() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Printf(c.Request.URL.Path)

		// c.Redirect(http.StatusOK, "http://localhost:50054")
		// c.Next()

		var proxyUrl = new(url.URL)
		proxyUrl.Scheme = "http"
		proxyUrl.Host = "localhost:50053"
		//u.Path = "base" // 这边若是赋值了，做转发的时候，会带上path前缀，例： /hello -> /base/hello

		proxy := httputil.NewSingleHostReverseProxy(proxyUrl)

		proxy.ServeHTTP(c.Writer, c.Request)

		c.Abort()
	}
}
