package routes

import (
	"koronet_web_app/handlers"

	"github.com/gin-gonic/gin"
)

type UserRoutes struct {
	userHandler *handlers.UserHandler
}

func NewUserRoutes(userHandler *handlers.UserHandler) *UserRoutes {
	return &UserRoutes{userHandler: userHandler}
}

func (ur *UserRoutes) GetRoutes(router *gin.Engine) {
	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/", ur.userHandler.CreateUser)
		userRoutes.GET("/:id", ur.userHandler.GetUser)
		userRoutes.GET("/", ur.userHandler.GetAllUsers)
		userRoutes.DELETE("/:id", ur.userHandler.DeleteUser)
	}
}
