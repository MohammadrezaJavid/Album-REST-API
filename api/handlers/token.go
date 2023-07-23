package handlers

import (
	"album/authentication"
	"album/database/models"
	"album/database/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// TODO swagger notations
// generate one token
func GenerateToken(ctx *gin.Context) {
	var request *TokenRequest
	var user *models.User

	// bind data of user from json to user of models.User
	if err := ctx.BindJSON(&request); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest,
			gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	// check if email exists and password is correct

	row := repositories.DataBase.Where("email = ?", request.Email).First(&user)
	if row.Error != nil {
		ctx.IndentedJSON(http.StatusInternalServerError,
			gin.H{"error": row.Error.Error()})
		ctx.Abort()
		return
	}

	credentialError := user.CheckPassword(request.Password)
	if credentialError != nil {
		ctx.IndentedJSON(http.StatusUnauthorized,
			gin.H{"error": "invalid credentials"})
		ctx.Abort()
		return
	}

	// create token
	tokenString, err := authentication.GenerateJwt(user.Email, user.Username)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	// send token by json file
	ctx.IndentedJSON(http.StatusOK, gin.H{"token": tokenString})
}
