package service

import (
	"album/infrastructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Hi(c *gin.Context) {
	infrastructure.LoadEnv()
	infrastructure.DatabaseHandle()
	c.IndentedJSON(http.StatusOK, gin.H{"message": "hello frinds"})
}
