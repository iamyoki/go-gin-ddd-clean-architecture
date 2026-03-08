package usecase

import (
	"context"
	"todo_api/module/todo/domain/todo"

	"github.com/google/uuid"
)

type DeleteTodo struct {
	Repo todo.Repository
}

func (usecase *DeleteTodo) Execute(ctx context.Context, id uuid.UUID) (*todo.Todo, error) {
	return usecase.Repo.DeleteById(ctx, id)
}
