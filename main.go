package main

import (
	"album/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	api.Routes(router)
	router.Run(":8070")
}
