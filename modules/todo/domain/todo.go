package domain

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	ID        uuid.UUID
	Title     string
	Completed bool
	CreatedAt time.Time
	UpdatedAt time.Time
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

func (todo *Todo) Complete() {
	todo.Completed = true
	todo.UpdatedAt = time.Now()
}
