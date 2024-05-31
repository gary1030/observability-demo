package handler

import (
	"net/http"

	"github.com/gary1030/learning-o11y/server/internal/repository/task"
	"github.com/gary1030/learning-o11y/server/pkg/log"
	"github.com/gary1030/learning-o11y/server/pkg/otel"
	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	taskRepository *task.Repository
}

func NewTaskHandler() *TaskHandler {
	repo, err := task.NewRepository()
	if err != nil {
		log.Fatal(err.Error())
	}

	return &TaskHandler{
		taskRepository: repo,
	}
}

func (t *TaskHandler) GetTasks(c *gin.Context) {
	ctx := c.Request.Context()
	_, span := otel.StartNewSpan(ctx)
	defer span.End()

	tasks, err := t.taskRepository.ListTasks(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching tasks"})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (t *TaskHandler) CreateTask(c *gin.Context) {
	ctx := c.Request.Context()
	_, span := otel.StartNewSpan(ctx)
	defer span.End()

	var input struct {
		Description string `json:"description"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task data"})
		return
	}

	if input.Description == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Description cannot be empty"})
		return
	}

	task, err := t.taskRepository.CreateTask(ctx, input.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Task created successfully", "data": task})
}
