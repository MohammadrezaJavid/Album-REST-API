package services

import (
	"album/database/models"

	"gorm.io/gorm"
)

func SelectAlbums(db *gorm.DB) *[]models.Album {
	var albums *[]models.Album
	db.Find(&albums)
	return albums
}

func SelectAlbumByID(ID string, db *gorm.DB) *models.Album {
	var album *models.Album
	db.Where("id = ?", ID).First(&album)
	return album
}

/*
*
curl -X POST -H "Content-Type: application/json" -d '{"Title": "javid photo", "Artist": "javid", "Price": 99.9}' http://localhost:8070/api/albums
*
*/

func InsertAlbum(newAlbum *models.Album, db *gorm.DB) error {
	result := db.Create(&newAlbum)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func UpdateAlbum(album *models.Album, db *gorm.DB) (string, error) {
	DeleteAlbumByID(album.ID, db)
	err := InsertAlbum(album, db)
	return album.ID, err
}

func DeleteAlbumByID(ID string, db *gorm.DB) *gorm.DB {
	result := db.Delete(&models.Album{}, ID)
	return result
}
