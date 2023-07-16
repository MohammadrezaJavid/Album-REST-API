package api

import (
	"album/api/handlers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.GET("/api/albums", handlers.GetAlbums)           // Get all album
	router.POST("/api/albums", handlers.PostAlbum)          // Add one album
	router.GET("/api/albums/:id", handlers.GetAlbumByID)    // Get one album by ID
	router.PUT("/api/albums/:id", handlers.UpdateAlbumByID) // Update one album by ID
	router.DELETE("/api/albums/:id", handlers.DeleteAlbum)  // Delete one album by ID
}
