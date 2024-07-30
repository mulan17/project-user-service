package routes

import (
	"net/http"
	"sign/utils"

	"github.com/gin-gonic/gin"
)

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Problem with generating a token"})
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login succesful", "token": token})
	
}