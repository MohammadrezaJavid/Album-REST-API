package main

import (
	"album/api"
	"album/config"
	"album/database"
	"album/database/repositories"
	"fmt"

	_ "album/docs"

	"github.com/gin-gonic/gin"
)

// @title Go + Gin Album rest API Service
// @version 1.0
// @description This is a sample server for save Albums in mysql database

// @contact.name API Support
// @contact.url http://www.swagger.io/api
// @contact.email api@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8070
// @BasePath /api
// @query.collection.format multi

func main() {
	setConfig()
	router := gin.Default()
	api.Routes(router)
	router.Run(":8080")
}

func setConfig() {
	config.LoadEnv()
	db := repositories.DatabaseHandle()
	if err := database.MigrateUp_1(db); err != nil {
		fmt.Printf("Error Migration: %s\n", err.Error())
	}
}
