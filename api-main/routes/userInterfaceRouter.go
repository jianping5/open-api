package routes

import (
	"api-main/api"

	"github.com/gin-gonic/gin"
)

func addUserInterfaceRouter(apiGroup *gin.RouterGroup) {
	userInterfaceRouter := apiGroup.Group("userInterface")
	{
		/*
			top		获取调用次数最高的接口
		*/
		userInterfaceRouter.GET("top", api.ListTopInvokeInterfaceInfo)
	}
}