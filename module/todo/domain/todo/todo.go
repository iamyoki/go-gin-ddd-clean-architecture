package todo

import (
	"time"

	"github.com/iamyoki/go-gin-ddd-clean-architecture/app/apperror"

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
	now := time.Now()
	return &Todo{
		ID:        uuid.New(),
		Title:     title,
		Completed: false,
		CreatedAt: now,
		UpdatedAt: now,
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
