package usecase

import (
	"context"
	"todo_api/module/todo/domain/todo"
)

type GetAllTodos struct {
	repo todo.Repository
}

func NewGetAllTodosUseCase(repo todo.Repository) *GetAllTodos {
	return &GetAllTodos{repo: repo}
}

func (usecase *GetAllTodos) Execute(ctx context.Context) ([]todo.Todo, error) {
	todos, err := usecase.repo.FindAll(ctx)

	if err != nil {
		return nil, err
	}

	return todos, nil
}
