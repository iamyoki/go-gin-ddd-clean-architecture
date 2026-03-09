package usecase

import (
	"context"

	"github.com/iamyoki/go-gin-ddd-clean-architecture/module/todo/domain/todo"

	"github.com/google/uuid"
)

type GetTodo struct {
	Repo todo.Repository
}

func (usecase *GetTodo) Execute(ctx context.Context, id uuid.UUID) (*todo.Todo, error) {
	return usecase.Repo.FindById(ctx, id)
}
