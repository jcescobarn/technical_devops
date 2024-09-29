package handlers

import (
	"koronet_web_app/repositories"
	"koronet_web_app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"koronet_web_app/entities"
)

type UserRequestBody struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserHandler struct {
	userRepository *repositories.UserRepository
	utilsFunctions *utils.Functions
}

func NewUserHandler(userRepository *repositories.UserRepository, utilsFunctions *utils.Functions) *UserHandler {
	return &UserHandler{userRepository: userRepository, utilsFunctions: utilsFunctions}
}

func (uh *UserHandler) CreateUser(context *gin.Context) {
	var requestBody UserRequestBody
	var err error

	if err = context.ShouldBindJSON(&requestBody); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := uh.utilsFunctions.HashPassword(requestBody.Password)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}

	user := &entities.User{
		Username: requestBody.Username,
		Name:     requestBody.Name,
		Email:    requestBody.Email,
		Password: hashedPassword,
	}

	if err = uh.userRepository.Create(user); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
		"user":    user.Username,
	})
}

func (uh *UserHandler) DeleteUser(context *gin.Context) {
	var userID string = context.Param("id")

	if err := uh.userRepository.Delete(userID); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func (uh *UserHandler) GetUser(context *gin.Context) {
	var userID string = context.Param("id")

	user, err := uh.userRepository.Get(userID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
		return
	}

	if user == nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	context.JSON(http.StatusOK, user)
}

func (uh *UserHandler) GetAllUsers(context *gin.Context) {
	users, err := uh.userRepository.GetAll()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	context.JSON(http.StatusOK, users)
}
