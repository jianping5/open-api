package routes

import (
	"api-main/middlewares"
	"api-main/models"
	"encoding/gob"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"

	// "github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

var R *gin.Engine

// 初始化 Gin
func init() {
	// 注册 User，否则他不能添加到 session 中（参与编解码需要先注册）
	gob.Register(models.User{})
	r := gin.Default()
	// 创建基于 cookie 的存储引擎，open-api 参数是用于加密的密钥
	store := cookie.NewStore([]byte("open-api"))
	// 设置 session 中间件，参数 mysession，指的是 session 的名字
	r.Use(middlewares.Cors(), sessions.Sessions("mysession", store))

	// todo：添加路由分组
	apiGroup := r.Group("/")
	addUserRouter(apiGroup)
	addInterfaceRouter(apiGroup)
	addUserInterfaceRouter(apiGroup)
	R = r
}
