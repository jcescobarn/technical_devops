package handlers

import (
	"koronet_web_app/repositories"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type SessionRequestBody struct {
	SessionID string `json:"session_id" binding:"required"`
	Data      string `json:"data" binding:"required"`
	TTL       int    `json:"ttl" binding:"required"` // Tiempo de expiración en segundos
}

type SessionHandler struct {
	sessionRepository *repositories.SessionRepository
}

func NewSessionHandler(sessionRepository *repositories.SessionRepository) *SessionHandler {
	return &SessionHandler{sessionRepository: sessionRepository}
}

func (sh *SessionHandler) CreateSession(context *gin.Context) {
	var request_body SessionRequestBody
	var err error

	if err = context.ShouldBindJSON(&request_body); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = sh.sessionRepository.Create(request_body.SessionID, request_body.Data, time.Duration(request_body.TTL)*time.Second)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create session"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Session created successfully"})
}

func (sh *SessionHandler) GetSession(context *gin.Context) {
	sessionID := context.Param("session_id")

	data, err := sh.sessionRepository.Get(sessionID)
	if err != nil {
		if err == redis.Nil {
			context.JSON(http.StatusNotFound, gin.H{"error": "Session not found"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get session"})
		}
		return
	}

	context.JSON(http.StatusOK, gin.H{"session_id": sessionID, "data": data})
}

func (sh *SessionHandler) DeleteSession(context *gin.Context) {
	sessionID := context.Param("session_id")

	err := sh.sessionRepository.Delete(sessionID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete session"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Session deleted successfully"})
}

func (sh *SessionHandler) UpdateSession(context *gin.Context) {
	var request_body SessionRequestBody
	var err error

	if err = context.ShouldBindJSON(&request_body); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = sh.sessionRepository.Update(request_body.SessionID, request_body.Data, time.Duration(request_body.TTL)*time.Second)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update session"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Session updated successfully"})
}
