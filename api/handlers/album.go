package handlers

// swag init --parseDependency --parseInternal

import (
	"album/database/services"

	"net/http"
	"strconv"

	"album/database/models"

	"github.com/gin-gonic/gin"
)

// @Summary 	Get all Albums from database
// @Description Retrieve all albums from the database
// @Tags		CRUD Album
// @Produce 	json
// @Success 	200 	 {array} 	models.Album
// @Router 		/jwt/albums  	[get]
// @Security BearerAuth
func GetAlbums(ctx *gin.Context) {
	albums := services.SelectAlbums()
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

	uintId, shouldReturn := strToInt(ID, ctx)
	if shouldReturn {
		return
	}

	album := services.SelectAlbumByID(uintId)
	if album.ID == 0 {
		ctx.IndentedJSON(http.StatusOK, gin.H{"message": "album not found"})
		return
	}
	ctx.JSON(http.StatusOK, album)
}

// PostAlbum	godoc
// @Summary 	Post an Album
// @Description Post an Album and saved to database
// @Tags		CRUD Album
// @Produce      json
// @Param		album	body		models.SwagAlbum	true	"Album JSON"
// @Success 	200 	{object} 	models.Album
// @Router 		/jwt/albums 	[post]
// @Security BearerAuth
func PostAlbum(ctx *gin.Context) {
	var newAlbum *models.Album
	if err := ctx.BindJSON(&newAlbum); err != nil {
		ctx.IndentedJSON(http.StatusFound, gin.H{"message": err.Error()})
		return
	}

	err := services.InsertAlbum(newAlbum)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusCreated, newAlbum)
}

// @Summary 	Update one Album from database
// @Description Retrieve an Album by id from database, update it and save to database
// @Tags		CRUD Album
// @Produce 	json
// @Param		id	path string true "Album ID"
// @Param		album	body		models.SwagAlbum	true	"Album JSON"
// @Success 	200 	{object} 	models.Album
// @Router 		/jwt/albums/{id}	[put]
// @Security BearerAuth
func UpdateAlbumByID(ctx *gin.Context) {

	// get ID
	ID := ctx.Param("id")
	uintId, shouldReturn := strToInt(ID, ctx)
	if shouldReturn {
		return
	}

	// get new album
	var newAlbum *models.SwagAlbum
	err := ctx.BindJSON(&newAlbum)
	if err != nil {
		ctx.IndentedJSON(http.StatusFound, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	// update old album
	if err := services.UpdateAlbum(uintId, newAlbum); err != nil {
		ctx.IndentedJSON(http.StatusOK, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	album := services.SelectAlbumByID(uintId)
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
	uintId, shouldReturn := strToInt(ID, ctx)
	if shouldReturn {
		return
	}

	result := services.DeleteAlbumByID(uintId)
	if result.Error != nil || result.RowsAffected == 0 {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Album was not deleted or Not found"})
	} else {
		ctx.IndentedJSON(http.StatusOK, gin.H{"message": "successfully deleted album"})
	}
}

func strToInt(ID string, ctx *gin.Context) (uint, bool) {
	uintId, err := strconv.ParseUint(ID, 10, 64)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter"})
		ctx.Abort()
		return 0, true
	}
	return uint(uintId), false
}
