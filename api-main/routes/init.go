package routes

import "github.com/gin-gonic/gin"

var R *gin.Engine

// 初始化 Gin
func init() {
	r := gin.Default()
	// todo：全局中间件并添加路由分组
	apiGroup := r.Group("main")
	addUserRouter(apiGroup)
	R = r
}