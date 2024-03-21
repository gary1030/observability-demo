package handler

import (
	"math/rand"
	"net/http"
	"time"

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
	c.JSON(status[n], "world")
}
