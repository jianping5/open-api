package main

import (
	"api-main/config"
	"api-main/routes"
)

func main() {
	err := routes.R.Run(config.Config.Gin.Address)
	if err != nil {
		return
	}
}