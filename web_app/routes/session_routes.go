package routes

import (
	"koronet_web_app/handlers"

	"github.com/gin-gonic/gin"
)

type SessionRoutes struct {
	sessionHandler *handlers.SessionHandler
}

func NewSessionRoutes(sessionHandler *handlers.SessionHandler) *SessionRoutes {
	return &SessionRoutes{sessionHandler: sessionHandler}
}

func (sr *SessionRoutes) GetRoutes(router *gin.Engine) {
	sessionRoutes := router.Group("/session")
	sessionRoutes.POST("/", sr.sessionHandler.CreateSession)
	sessionRoutes.GET("/:session_id", sr.sessionHandler.GetSession)
	sessionRoutes.DELETE("/:session_id", sr.sessionHandler.DeleteSession)
	sessionRoutes.PUT("/", sr.sessionHandler.UpdateSession)
}
