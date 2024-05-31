package handler

import (
	"net/http"

	"github.com/gary1030/learning-o11y/server/internal/repository/joke"
	"github.com/gary1030/learning-o11y/server/pkg/log"
	"github.com/gin-gonic/gin"
)

type JokeHandler struct {
	jokeRepository *joke.Repository
}

func NewJokeHandler() *JokeHandler {
	repo, err := joke.NewRepository()
	if err != nil {
		log.Fatal(err.Error())
	}

	return &JokeHandler{
		jokeRepository: repo,
	}
}

func (j *JokeHandler) GetJoke(c *gin.Context) {
	ctx := c.Request.Context()

	joke, err := j.jokeRepository.GetRandomJoke(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching joke"})
		return
	}
	c.JSON(http.StatusOK, joke)
}
