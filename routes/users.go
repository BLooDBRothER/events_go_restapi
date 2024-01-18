package routes

import (
	"net/http"

	"example.com/events/controller"
	"example.com/events/jwt"
	"example.com/events/models"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	userId, err := controller.CreateUser(&user);

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create user. Try again later."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"id": userId})
}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	isValidUser, err := controller.ValidateUser(&user)

	if(!isValidUser) {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid email or password"})
		return
	}

	jwtToken, err := jwt.GenerateJwtToken(user.Id, user.Email)

	if(err != nil) {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authorize please try again later"})
	}

	context.JSON(http.StatusOK, gin.H{"token": jwtToken})
}

