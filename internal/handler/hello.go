package handler

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gary1030/learning-o11y/pkg/otel"
	"github.com/gin-gonic/gin"
)

func HelloHandler(c *gin.Context) {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	status := []int{
		http.StatusOK,
		http.StatusBadRequest,
		http.StatusUnauthorized,
		http.StatusNotFound,
		http.StatusInternalServerError,
	}
	n := rng.Int() % len(status)

	tempStall(c)
	tempStall(c)
	c.JSON(status[n], "world")
}

func tempStall(c *gin.Context) {
	ctx := c.Request.Context()
	_, span := otel.StartNewSpan(ctx)
	defer span.End()

	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	delay := rng.Intn(300)
	time.Sleep(time.Duration(delay) * time.Millisecond)
}
