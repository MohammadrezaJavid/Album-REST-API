package main

import (
	"album/api"
	"album/config"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	router := gin.Default()
	api.Routes(router)
	router.Run(":8070")
}
