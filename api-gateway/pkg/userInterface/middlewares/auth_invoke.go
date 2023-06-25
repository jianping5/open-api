package middlewares

import (
	"api-gateway/pkg/userInterface/inner"
	"api-gateway/pkg/userInterface"
	"fmt"

	"github.com/gin-gonic/gin"
)

func AuthInvoke() gin.HandlerFunc {
	return func(c *gin.Context) {

		Svc := userInterface.GetSvc()

		// 认证并增加接口调用次数
		inner.AddUserInterfaceInfo(c, Svc.Client)

		fmt.Println("AuthInvoke() 执行")

		// if !res {
		// 	c.Abort()
		// }
		c.Next()
	}
}
