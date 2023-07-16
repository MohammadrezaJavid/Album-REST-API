package main

import (
	"album/api"
	"album/database"
	"album/database/repositories"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	setConfig()
	router := gin.Default()
	api.Routes(router)
	router.Run(":8070")
}

func setConfig() {
	// config.LoadEnv()

	db := repositories.DatabaseHandle()
	if err := database.MigrateUp_1(db); err != nil {
		fmt.Printf("Error Migration: %s\n", err.Error())
	}
}
