package main

import (
	"album/api"
	"album/config"
	"album/database"
	"album/database/repositories"
	"fmt"

	_ "album/docs"
)

func init() {
	fmt.Println(">>Initializing config...<<")
	config.LoadEnv()
	db := repositories.DatabaseHandle()
	if err := database.MigrateUp_1(db); err != nil {
		fmt.Printf("Error Migration: %s\n", err.Error())
	}
}

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
	router := api.Routes()
	router.Run(":8070")
}
