package routes

import (
	"api-main/api"

	"github.com/gin-gonic/gin"
)

func addInterfaceRouter(apiGroup *gin.RouterGroup) {
	interfaceRouter := apiGroup.Group("interfaceInfo")
	{
		/*
			add		创建接口
			delete  删除接口
			update  修改接口
			get     获取接口
			online  上线接口
			offline 下线接口
			invoke  调用接口
			list/page 分页获取接口列表
		*/
		interfaceRouter.POST("add", api.AddInterfaceInfo)
		interfaceRouter.POST("delete", api.DeleteInterfaceInfo)
		interfaceRouter.POST("update", api.UpdateInterfaceInfo)
		interfaceRouter.GET("get", api.GetInterfaceInfoById)
		interfaceRouter.POST("online", api.OnlineInterfaceInfo)
		interfaceRouter.POST("offline", api.OffLineInterfaceInfo)
		interfaceRouter.POST("invoke", api.InvokeInterfaceInfo)
		interfaceRouter.GET("list/page", api.ListInterfaceInfoByPage)
	}
}
