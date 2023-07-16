package database

import (
	"album/database/models"

	"gorm.io/gorm"
)

func MigrateUp_1(ctx *gorm.DB) error {
	err := ctx.AutoMigrate(&models.Album{})
	if err != nil {
		return err
	}
	return nil
}

func MigrateDown_1(ctx *gorm.DB) error {
	err := ctx.Migrator().DropTable(&models.Album{})
	if err != nil {
		return err
	}
	return nil
}
