package services

import (
	"album/database/models"

	"gorm.io/gorm"
)

func InsertUser(newUser *models.User, db *gorm.DB) error {
	row := db.Create(&newUser)
	if row.Error != nil {
		return row.Error
	}
	return nil
}
