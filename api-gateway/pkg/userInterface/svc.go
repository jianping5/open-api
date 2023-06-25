package userInterface

import (
	// "api-gateway/pkg/auth"
	"api-gateway/pkg/config"
	// "api-gateway/pkg/userInterface/routes"

	"github.com/gin-gonic/gin"
)

var Svc *ServiceClient

func InitSvc(r *gin.Engine, c *config.Config) {
	Svc = &ServiceClient{
		Client: InitServiceClient(c),
	}

	// routes := r.Group("api")
	// routes.POST("", svc.AddUserInterfaceInfo)
	// routes.GET("hello", svc.AddUserInterfaceInfo)
}

func GetSvc() *ServiceClient {
	return Svc
}

// func (svc *ServiceClient) AddUserInterfaceInfo(ctx *gin.Context) {
// 	routes.AddUserInterfaceInfo(ctx, svc.Client)
// }
