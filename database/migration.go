package database

import (
	"album/database/models"
	"album/database/repositories"
	"errors"
	"log"
)

func Migrate() {

	if err := repositories.DataBase.AutoMigrate(&models.Album{}); err != nil {
		log.Fatal(errors.New(err.Error()))
	}

	if err := repositories.DataBase.AutoMigrate(&models.User{}); err != nil {
		log.Fatal(errors.New(err.Error()))
	}
}

func MigrateDown() {

	if err := repositories.DataBase.Migrator().DropTable(&models.Album{}); err != nil {
		log.Fatal(errors.New(err.Error()))
	}

	if err := repositories.DataBase.Migrator().DropTable(&models.User{}); err != nil {
		log.Fatal(errors.New(err.Error()))
	}
}
