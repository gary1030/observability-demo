package routes

import (
	"github.com/gary1030/learning-o11y/server/internal/handler"
	"github.com/gin-gonic/gin"
)

func SetTaskRoute(e *gin.Engine) {
	taskHandler := handler.NewTaskHandler()
	taskGroup := e.Group("/task")
	{
		taskGroup.GET("", taskHandler.GetTasks)
		taskGroup.POST("", taskHandler.CreateTask)
	}
}
