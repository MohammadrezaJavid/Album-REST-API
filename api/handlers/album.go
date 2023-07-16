package handlers

import (
	"album/database/repositories"
	"album/database/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB *gorm.DB = repositories.DatabaseHandle()

func GetAlbums(ctx *gin.Context) {
	albums := services.SelectAlbums(DB)
	ctx.IndentedJSON(http.StatusOK, albums)
}

func GetAlbumByID(ctx *gin.Context) {
	ID := ctx.Param("id")
	album := services.SelectAlbumByID(ID, DB)
	ctx.IndentedJSON(http.StatusOK, album)
}

func PostAlbum(ctx *gin.Context) {
	result, ID := services.InsertAlbum(ctx, DB)
	if err := result.Error; err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
	}
	if ID > -1 {
		ctx.IndentedJSON(http.StatusOK, gin.H{"Album added by ID": ID})
	}
}

func UpdateAlbumByID(ctx *gin.Context) {
	// TODO
}

func DeleteAlbum(ctx *gin.Context) {
	ID := ctx.Param("id")
	services.DeleteAlbumByID(ID, DB)
}
