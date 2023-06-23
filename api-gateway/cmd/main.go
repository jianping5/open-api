package main

import (
	"api-gateway/pkg/config"
	"api-gateway/pkg/interfaces"
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
	r.Use(middlewares.Conv())

	interfaces.RegisterRoutes(r, &c)

	r.Run(c.Port)
}
