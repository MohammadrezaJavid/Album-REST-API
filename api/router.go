package api

import (
	"album/api/handlers"
	"album/middlewares"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Routes() *gin.Engine {

	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		api.POST("/token", handlers.GenerateToken)
		api.POST("/user/register", handlers.RegisterUser)

		jwt := api.Group("/jwt").Use(middlewares.Authorization())
		{
			jwt.GET("/albums", handlers.GetAlbums)          // Get all album
			jwt.POST("/albums", handlers.PostAlbum)         // Add one album
			jwt.PUT("/albums", handlers.UpdateAlbumByID)    // Update one album by ID
			jwt.GET("/albums/:id", handlers.GetAlbumByID)   // Get one album by ID
			jwt.DELETE("/albums/:id", handlers.DeleteAlbum) // Delete one album by ID

		}
	}
	return router
}
