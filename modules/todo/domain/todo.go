package domain

import (
	"time"
	apperror "todo_api/app/error"

	"github.com/google/uuid"
)

type Todo struct {
	ID          uuid.UUID
	Title       string
	Completed   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	CompletedAt *time.Time
}

func Create(title string) *Todo {
	return &Todo{
		ID:        uuid.New(),
		Title:     title,
		Completed: false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (todo *Todo) Complete() error {
	if todo.Completed {
		return &apperror.BadRequest{Msg: "todo already completed"}
	}

	todo.Completed = true
	now := time.Now()
	todo.UpdatedAt = now
	todo.CompletedAt = &now
	return nil
}
