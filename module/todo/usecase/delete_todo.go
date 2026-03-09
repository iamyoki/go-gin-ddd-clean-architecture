package usecase

import (
	"context"

	"github.com/iamyoki/go-gin-ddd-clean-architecture/module/todo/domain/todo"

	"github.com/google/uuid"
)

type DeleteTodo struct {
	Repo todo.Repository
}

func (usecase *DeleteTodo) Execute(ctx context.Context, id uuid.UUID) (*todo.Todo, error) {
	return usecase.Repo.DeleteById(ctx, id)
}
