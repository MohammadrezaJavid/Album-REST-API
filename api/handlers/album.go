package handlers

import (
	"album/database/models"
	"album/database/repositories"
	"album/database/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary 	Get all Albums from database
// @Description Retrieve all albums from the database
// @Tags		get album
// @Produce 	json
// @Success 	200 	 {object} 	models.Album
// @Router 		/albums  [get]
func GetAlbums(ctx *gin.Context) {
	albums := services.SelectAlbums(repositories.DataBase)
	ctx.IndentedJSON(http.StatusOK, albums)
}

// @Summary 	Get one Album from database
// @Description Retrieve an Album by id from the database
// @Tags		get album
// @Produce 	json
// @Param		id	path string true "Album ID"
// @Success 	200 {object} models.Album
// @Router 		/albums/{id}	[get]
func GetAlbumByID(ctx *gin.Context) {
	ID := ctx.Param("id")
	album := services.SelectAlbumByID(ID, repositories.DataBase)
	ctx.JSON(http.StatusOK, album)
}

// PostAlbum	godoc
// @Summary 	Post an Album
// @Description Post an Album and saved to database
// @Tags		post album
// @Produce      json
// @Param		album	body		models.Album	true	"Album JSON"
// @Success 	200 	{object} 	models.Album
// @Router 		/albums [post]
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
// @Tags		put album
// @Produce 	json
// @Param		album	body		models.Album	true	"Album JSON"
// @Success 	200 	{object} models.Album
// @Router 		/albums	[put]
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
// @Tags		delete album
// @Param		id	path	string	true	"Album id"
// @Success 	200 	{object} string
// @Router 		/albums/{id}	[delete]
func DeleteAlbum(ctx *gin.Context) {
	ID := ctx.Param("id")
	album := services.SelectAlbumByID(ID, repositories.DataBase)
	if album.ID == "" {
		err := message{err: "album not found"}
		ctx.IndentedJSON(http.StatusNotFound, err.err)
		return
	}

	services.DeleteAlbumByID(ID, repositories.DataBase)
	err := message{err: "successfully deleted album"}
	ctx.IndentedJSON(http.StatusOK, err.err)
}

type message struct {
	err string
}
