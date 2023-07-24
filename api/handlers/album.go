package handlers

import (
	"album/database/models"
	"album/database/repositories"
	"album/database/services"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary 	Get all Albums from database
// @Description Retrieve all albums from the database
// @Tags		CRUD Album
// @Produce 	json
// @Success 	200 	 {object} 	models.Album
// @Router 		/jwt/albums  	[get]
// @Security BearerAuth
func GetAlbums(ctx *gin.Context) {
	albums := services.SelectAlbums(repositories.DataBase)
	ctx.IndentedJSON(http.StatusOK, albums)
}

// @Summary 	Get one Album from database
// @Description Retrieve an Album by id from the database
// @Tags		CRUD Album
// @Produce 	json
// @Param		id	path string true "Album ID"
// @Success 	200 {object} models.Album
// @Router 		/jwt/albums/{id}	[get]
// @Security BearerAuth
func GetAlbumByID(ctx *gin.Context) {
	ID := ctx.Param("id")
	album := services.SelectAlbumByID(ID, repositories.DataBase)
	ctx.JSON(http.StatusOK, album)
}

// PostAlbum	godoc
// @Summary 	Post an Album
// @Description Post an Album and saved to database
// @Tags		CRUD Album
// @Produce      json
// @Param		album	body		models.Album	true	"Album JSON"
// @Success 	200 	{object} 	models.Album
// @Router 		/jwt/albums 	[post]
// @Security BearerAuth
func PostAlbum(ctx *gin.Context) {
	var newAlbum *models.Album
	if err := ctx.BindJSON(&newAlbum); err != nil {
		ctx.IndentedJSON(http.StatusFound, gin.H{"message": err.Error()})
		return
	}

	err := services.InsertAlbum(newAlbum, repositories.DataBase)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, newAlbum)
}

// @Summary 	Update one Album from database
// @Description Retrieve an Album by id from database, update it and save to database
// @Tags		CRUD Album
// @Produce 	json
// @Param		album	body		models.Album	true	"Album JSON"
// @Success 	200 	{object} models.Album
// @Router 		/jwt/albums		[put]
// @Security BearerAuth
func UpdateAlbumByID(ctx *gin.Context) {
	var updateAlbum *models.Album
	err := ctx.BindJSON(&updateAlbum)
	if err != nil {
		ctx.IndentedJSON(http.StatusFound, gin.H{"message": err.Error()})
		return
	}

	id, err := services.UpdateAlbum(updateAlbum, repositories.DataBase)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	album := services.SelectAlbumByID(id, repositories.DataBase)
	ctx.IndentedJSON(http.StatusOK, album)
}

// @Summary 	Delete one Album from database
// @Description Delete an Album by id from database.
// @Tags		CRUD Album
// @Param		id	path	string	true	"Album id"
// @Success 	200 	{object} string
// @Router 		/jwt/albums/{id}	[delete]
// @Security BearerAuth
func DeleteAlbum(ctx *gin.Context) {
	ID := ctx.Param("id")
	album := services.SelectAlbumByID(ID, repositories.DataBase)
	if album.ID == "" {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"Error": errors.New("album not found")})
		ctx.Abort()
		return
	}

	services.DeleteAlbumByID(ID, repositories.DataBase)
	ctx.IndentedJSON(http.StatusOK, gin.H{"Error": errors.New("successfully deleted album")})
}
