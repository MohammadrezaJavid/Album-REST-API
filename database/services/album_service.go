package services

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Album struct {
	ID     int64
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func SelectAlbums(db *gorm.DB) *[]Album {
	var albums *[]Album
	db.Find(&albums)
	return albums
}

func SelectAlbumByID(ID string, db *gorm.DB) *[]Album {
	var album *[]Album
	db.Find(&album, ID)
	return album
}

/*
*
curl -X POST -H "Content-Type: application/json" -d '{"Title": "javid photo", "Artist": "javid", "Price": 99.9}' http://localhost:8070/api/albums
*
*/
func InsertAlbum(ctx *gin.Context, db *gorm.DB) (*gorm.DB, int64) {
	var newAlbum *Album
	err := ctx.BindJSON(&newAlbum)
	if err != nil {
		fmt.Println(newAlbum)
		ctx.IndentedJSON(http.StatusFound, gin.H{"PostAlbumMessage": err.Error()})
		return nil, -1
	}

	result := db.Create(&newAlbum)
	return result, newAlbum.ID
}

func UpdateAlbumByID(ID string, db *gorm.DB) *Album {
	return nil
}

func DeleteAlbumByID(ID string, db *gorm.DB) {
	db.Delete(&Album{}, ID)
}
