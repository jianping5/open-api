package routes

import (
	"api-interface/api"

	"github.com/gin-gonic/gin"
)

func addDemoRouter(apiGroup *gin.RouterGroup) {
	demoRouter := apiGroup.Group("api")
	{
		/*
			hello
			name
		*/
		demoRouter.GET("hello", api.SayHello)
		demoRouter.POST("name", api.GetNameById)
	}
}
