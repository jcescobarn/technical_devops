package handlers

import (
	"koronet_web_app/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostRequestBody struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type PostHandler struct {
	postRepository *repositories.PostRepository
}

func NewPostHandler(postRepository *repositories.PostRepository) *PostHandler {
	return &PostHandler{postRepository: postRepository}
}

func (ph *PostHandler) CreatePost(context *gin.Context) {
	var requestBody PostRequestBody
	var err error

	if err = context.ShouldBindJSON(&requestBody); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post := repositories.Post{
		Title:   requestBody.Title,
		Content: requestBody.Content,
	}

	result, err := ph.postRepository.CreatePost(post)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Post created successfully", "post_id": result.InsertedID})
}

func (ph *PostHandler) GetPost(context *gin.Context) {
	id := context.Param("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	post, err := ph.postRepository.GetPost(objectID.Hex())
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch post"})
		return
	}

	if post == nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"post": post})
}

func (ph *PostHandler) GetAllPosts(context *gin.Context) {
	posts, err := ph.postRepository.GetAllPost()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch posts"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"posts": posts})
}

func (ph *PostHandler) DeletePost(context *gin.Context) {
	id := context.Param("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	result, err := ph.postRepository.DeletePost(objectID.Hex())
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete post"})
		return
	}

	if result.DeletedCount == 0 {
		context.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}
