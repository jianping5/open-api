package interfaces

import (
	// "api-gateway/pkg/auth"
	"api-gateway/pkg/config"
	"api-gateway/pkg/interfaces/routes"

	"github.com/gin-gonic/gin"
)

/* func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/order")
	routes.Use(a.AuthRequired)
	routes.POST("/", svc.CreateOrder)
} */
func RegisterRoutes(r *gin.Engine, c *config.Config) {
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/api")
	routes.POST("/", svc.GetName)
}

func (svc *ServiceClient) GetName(ctx *gin.Context) {
	routes.GetName(ctx, svc.Client)
}


