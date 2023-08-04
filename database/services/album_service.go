package services

import (
	"album/database/models"
	"album/database/repositories"
	"errors"

	"gorm.io/gorm"
)

func SelectAlbums() *[]models.Album {
	var albums *[]models.Album
	repositories.DataBase.Find(&albums)
	return albums
}

func SelectAlbumByID(ID uint) *models.Album {
	var album *models.Album
	repositories.DataBase.Where("id = ?", ID).First(&album)
	return album
}

/*
*
curl -X POST -H "Content-Type: application/json" -d '{"Title": "javid photo", "Artist": "javid", "Price": 99.9}' http://localhost:8070/api/albums
*
*/

func InsertAlbum(newAlbum *models.Album) error {
	result := repositories.DataBase.Create(&newAlbum)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func UpdateAlbum(ID uint, newAlbum *models.SwagAlbum) error {
	oldAlbum := SelectAlbumByID(ID)
	if oldAlbum.ID == 0 {
		return errors.New("album not found")
	}

	oldAlbum.Title = newAlbum.Title
	oldAlbum.Artist = newAlbum.Artist
	oldAlbum.Price = newAlbum.Price

	if err := repositories.DataBase.Save(&oldAlbum).Error; err != nil {
		return err
	}

	return nil
}

func DeleteAlbumByID(ID uint) *gorm.DB {
	result := repositories.DataBase.Delete(&models.Album{}, ID)
	return result
}
