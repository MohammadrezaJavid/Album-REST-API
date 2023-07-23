package handlers

import (
	"album/database/models"
	"album/database/repositories"
	"album/database/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Register one user
func RegisterUser(ctx *gin.Context) {
	// bind data of user from json to user of models.User
	var user models.User
	if err := ctx.BindJSON(&user); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest,
			gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	// hashed password
	if err := user.HashPassword(user.Password); err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	// insert user to database
	if err := services.InsertUser(&user, repositories.DataBase); err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	// output of a json; newly created user info
	ctx.IndentedJSON(
		http.StatusCreated,
		gin.H{
			"userId":   user.ID,
			"email":    user.Email,
			"username": user.Username,
		},
	)
}
