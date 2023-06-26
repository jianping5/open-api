package routes

import (
	"api-interface/middlewares"

	"github.com/gin-gonic/gin"
)

var R *gin.Engine

// 初始化 Gin
func init() {
	r := gin.Default()

	r.Use(middlewares.LimitSource())

	// todo：添加路由分组
	apiGroup := r.Group("/")
	addDemoRouter(apiGroup)
	R = r
}
