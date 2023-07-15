package api

import (
	"album/api/handlers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.GET("/api/albums", handlers.GetAlbums)
	router.POST("/api/albums", handlers.PostAlbum)
	router.GET("/api/albums/:id", handlers.GetAlbumByID)
	router.PUT("/api/albums/:id", handlers.UpdateAlbumByID)
	router.DELETE("/api/albums/:id", handlers.DeleteAlbum)
}
