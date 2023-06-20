package routes

import (
	"github.com/gin-gonic/gin"
	"api-main/api"
)

func addUserRouter(apiGroup *gin.RouterGroup) {
	userRouter := apiGroup.Group("user")
	{
		/*
			login		登录
			register	注册
			logout		登出
		*/
		userRouter.POST("register", api.UserRegister)
	}
}
