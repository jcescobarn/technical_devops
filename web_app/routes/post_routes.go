package routes

import (
	"koronet_web_app/handlers"

	"github.com/gin-gonic/gin"
)

type PostRoutes struct {
	postHandler *handlers.PostHandler
}

func NewPostRoutes(postHandler *handlers.PostHandler) *PostRoutes {
	return &PostRoutes{postHandler: postHandler}
}

func (pr *PostRoutes) GetRoutes(router *gin.Engine) {
	post_routes := router.Group("/posts")

	post_routes.GET("/", pr.postHandler.GetAllPosts)
	post_routes.GET("/:id", pr.postHandler.GetPost)
	post_routes.POST("/", pr.postHandler.CreatePost)
	post_routes.DELETE("/:id", pr.postHandler.DeletePost)
}
