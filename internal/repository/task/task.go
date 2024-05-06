package task

import (
	"context"

	"github.com/gary1030/learning-o11y/pkg/log"
	"github.com/gary1030/learning-o11y/pkg/otel"
	"github.com/gary1030/learning-o11y/pkg/sqlite"
)

type Task struct {
	ID          int
	Description string
	Done        bool
}

type Repository struct {
	db *sqlite.Database
}

func NewRepository() (*Repository, error) {
	db, err := sqlite.NewDatabase("task.sqlite.db")
	if err != nil {
		log.Fatal(err.Error())
	}

	const tableCreationQuery = `
	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		description TEXT NOT NULL,
		done BOOLEAN NOT NULL DEFAULT 0
	)`
	_, err = db.Exec(tableCreationQuery)
	if err != nil {
		return nil, err
	}

	return &Repository{db}, nil
}

func (r *Repository) CreateTask(ctx context.Context, description string) error {
	_, span := otel.StartNewSpan(ctx)
	defer span.End()

	query := `INSERT INTO tasks (description) VALUES (?)`
	_, err := r.db.Exec(query, description)
	return err
}

func (r *Repository) ListTasks(ctx context.Context) ([]*Task, error) {
	_, span := otel.StartNewSpan(ctx)
	defer span.End()

	rows, err := r.db.Query("SELECT id, description, done FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*Task
	for rows.Next() {
		var t Task
		if err := rows.Scan(&t.ID, &t.Description, &t.Done); err != nil {
			return nil, err
		}
		tasks = append(tasks, &t)
	}

	return tasks, nil
}
