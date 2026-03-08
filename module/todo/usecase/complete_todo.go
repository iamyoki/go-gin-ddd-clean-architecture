package usecase

import (
	"context"
	"todo_api/module/todo/domain/todo"

	"github.com/google/uuid"
)

type CompleteTodo struct {
	Repo todo.Repository
}

func (usecase *CompleteTodo) Execute(ctx context.Context, id uuid.UUID) (*todo.Todo, error) {
	todo, err := usecase.Repo.FindById(ctx, id)

	if err != nil {
		return nil, err
	}

	err = todo.Complete()

	if err != nil {
		return nil, err
	}

	err = usecase.Repo.Save(ctx, todo)

	if err != nil {
		return nil, err
	}

	return todo, nil
}
