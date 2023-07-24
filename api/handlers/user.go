package handlers

import (
	dbModel "album/database/models"
	"album/database/repositories"
	"album/database/services"
	swagModel "album/swagger/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterUser godoc
// @Summary		Register one user
// @Description Get info of one user and saved to database
// @Tags		Auth
// @Produce		json
// @Param		user	body		swagModel.User		true	"User JSON"
// @Success		200		{object}	swagModel.Result
// @Router		/user/register		[post]
func RegisterUser(ctx *gin.Context) {
	// bind data of user from json to user of models.User
	var user dbModel.User
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

		&swagModel.Result{
			UserId:   user.ID,
			Email:    user.Email,
			Username: user.Username,
		},
	)
}
