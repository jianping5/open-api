package routes

import (
	"github.com/gin-gonic/gin"
	// "api-interface/middlewares"
)

var R *gin.Engine

// 初始化 Gin
func init() {
	r := gin.Default()

	// r.Use(middlewares.Cors())
	
	// todo：添加路由分组
	apiGroup := r.Group("/")
	addDemoRouter(apiGroup)
	R = r
}