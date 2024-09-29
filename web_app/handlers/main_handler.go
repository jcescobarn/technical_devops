package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GreetingResponse struct {
	Greeting string
}

type MainHandler struct {
}

func NewMainHandler() *MainHandler {
	return &MainHandler{}
}

func (mh *MainHandler) Greeting(context *gin.Context) {

	context.JSON(http.StatusOK, GreetingResponse{
		Greeting: "Hi Koronet Team",
	})
	return

}
