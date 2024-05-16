package task

import (
	"context"
	"time"

	"github.com/gary1030/learning-o11y/pkg/database"
	"github.com/gary1030/learning-o11y/pkg/log"
	"github.com/gary1030/learning-o11y/pkg/otel"
)

type Task struct {
	ID          uint `gorm:"autoIncrement;primaryKey"`
	Description string
	Done        bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Repository struct {
	db *database.Database
}

func NewRepository() (*Repository, error) {
	db, err := database.NewDatabase("task.sqlite.db")
	if err != nil {
		log.Fatal(err.Error())
	}

	err = db.AutoMigrate(&Task{})
	if err != nil {
		return nil, err
	}

	return &Repository{db}, nil
}

func (r *Repository) CreateTask(ctx context.Context, description string) (*Task, error) {
	_, span := otel.StartNewSpan(ctx)
	defer span.End()

	task := Task{
		Description: description,
		Done:        false,
	}

	result := r.db.Create(&task)
	if result.Error != nil {
		return nil, result.Error
	}

	return &task, nil
}

func (r *Repository) ListTasks(ctx context.Context) ([]*Task, error) {
	_, span := otel.StartNewSpan(ctx)
	defer span.End()

	var tasks []*Task
	result := r.db.Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}

	return tasks, nil
}
