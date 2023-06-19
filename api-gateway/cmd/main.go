package main

import (
	"log"
	"api-gateway/pkg/config"
	"api-gateway/pkg/interfaces"
	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	interfaces.RegisterRoutes(r, &c)

	r.Run(c.Port)
}


