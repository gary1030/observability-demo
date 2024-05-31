package routes

import (
	"github.com/gary1030/learning-o11y/server/internal/handler"
	"github.com/gin-gonic/gin"
)

func SetJokeRoute(e *gin.Engine) {
	jokeHandler := handler.NewJokeHandler()
	jokeGroup := e.Group("/joke")
	{
		jokeGroup.GET("", jokeHandler.GetJoke)
	}
}
