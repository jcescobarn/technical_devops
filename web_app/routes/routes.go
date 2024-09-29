package routes

import (
	"koronet_web_app/handlers"

	"github.com/gin-gonic/gin"
)

type GeneralRoutes struct {
	MainHandler *handlers.MainHandler
}

func NewGeneralRoutes(mainHandler *handlers.MainHandler) *GeneralRoutes {
	return &GeneralRoutes{}
}

func (gr *GeneralRoutes) GetRoutes(router *gin.Engine) {
	general_routes := router.Group("/greeting")
	general_routes.GET("/", gr.MainHandler.Greeting)
}
