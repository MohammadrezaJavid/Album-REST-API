package main

import (
	"album/service"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/hi", service.Hi)
	router.Run(":80")
}
