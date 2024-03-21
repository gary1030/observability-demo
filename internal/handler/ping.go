package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "pong!")
}
