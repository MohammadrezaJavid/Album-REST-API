package database

import (
	"album/database/models"
	"album/database/repositories"
	"log"
)

func Migrate() {

	if err := repositories.DataBase.AutoMigrate(&models.Album{}); err != nil {
		log.Fatal(err.Error())
	}

	if err := repositories.DataBase.AutoMigrate(&models.User{}); err != nil {
		log.Fatal(err.Error())
	}
}

func MigrateDown() error {
	err := repositories.DataBase.Migrator().DropTable(&models.Album{})
	if err != nil {
		return err
	}
	return nil
}
