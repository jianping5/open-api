package routes

import (
	"api-main/api"

	"github.com/gin-gonic/gin"
)

func addUserRouter(apiGroup *gin.RouterGroup) {
	userRouter := apiGroup.Group("user")
	{
		/*
			login		登录
			register	注册
			logout		登出
		*/
		userRouter.POST("login", api.UserLogin)
		userRouter.POST("register", api.UserRegister)
		userRouter.POST("logout", api.UserLogout)
		userRouter.GET("get/login", api.GetCurrentUser)
	}
}
