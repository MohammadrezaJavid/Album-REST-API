package api

import (
	"album/api/handlers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Routes() *gin.Engine {

	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/albums", handlers.GetAlbums)          // Get all album
		api.POST("/albums", handlers.PostAlbum)         // Add one album
		api.PUT("/albums", handlers.UpdateAlbumByID)    // Update one album by ID
		api.GET("/albums/:id", handlers.GetAlbumByID)   // Get one album by ID
		api.DELETE("/albums/:id", handlers.DeleteAlbum) // Delete one album by ID
		api.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	return router
}
