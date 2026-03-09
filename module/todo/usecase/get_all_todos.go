package usecase

import (
	"context"
	"todo_api/module/todo/domain/todo"
)

type GetAllTodos struct {
	Repo todo.Repository
}

func (usecase *GetAllTodos) Execute(ctx context.Context) ([]todo.Todo, error) {
	todos, err := usecase.Repo.FindAll(ctx)

	if err != nil {
		return nil, err
	}

	return todos, nil
}
