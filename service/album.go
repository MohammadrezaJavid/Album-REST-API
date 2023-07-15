package service

import (
	"album/database"
	"net/http"

	"github.com/docker/docker/daemon/config"
	"github.com/gin-gonic/gin"
)

func Hi(c *gin.Context) {
	// infrastructure.LoadEnv()
	db := config.DatabaseHandle()
	// database.Down_1(db)
	database.Up_1(db)
	// infrastructure.MakeMigrate()

	c.IndentedJSON(http.StatusOK, gin.H{"message": "hello frinds"})
}
