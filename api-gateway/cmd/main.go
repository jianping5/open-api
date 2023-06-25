package main

import (
	"api-gateway/pkg/config"
	"api-gateway/pkg/interfaces"
	"api-gateway/pkg/userInterface"
	userInterface_middlewares "api-gateway/pkg/userInterface/middlewares"
	"api-gateway/pkg/middlewares"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()
	
	r.Use(userInterface_middlewares.AuthInvoke(), middlewares.Conv())

	interfaces.RegisterRoutes(r, &c)
	// 初始化 serviceClient
	userInterface.InitSvc(r, &c)

	r.Run(c.Port)
}
