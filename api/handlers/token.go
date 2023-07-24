package handlers

import (
	"album/authentication"
	dbModels "album/database/models"
	"album/database/repositories"
	swagModels "album/swagger/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GenerateToken godoc
// @Summary		Generate one token
// @Description Create one jwt token
// @Tags		Auth
// @Produce		json
// @Param		baseAuth	body		swagModels.BaseAuth	true	"baseAuth JSON"
// @Success		200			{object}	swagModels.Token
// @Router		/token		[post]
func GenerateToken(ctx *gin.Context) {
	var request *swagModels.BaseAuth
	var user dbModels.User

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
	ctx.IndentedJSON(http.StatusOK, &swagModels.Token{Token: tokenString})
}
