package routes

import (
	"net/http"

	"example.com/rest-api/models"
	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could notsave user."})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}


func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}


	err = user.ValidateCredential()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate the user." })
		return
	}
 

	 token, err := utils.GenerateToken(user.Email, int64(user.ID))
	 if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate the user." })
		return
	}
 
	context.JSON(http.StatusCreated, gin.H{"message": "Login Successfull.", "Token": token})
}

func getAllUsers(context *gin.Context) {
	users, err := models.GetAllUsers()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get users"})
		return
	}

	context.JSON(http.StatusOK, users)
}
