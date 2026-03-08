package usecase

import (
	"context"
	"todo_api/module/todo/domain/todo"
)

type CreateTodo struct {
	repo todo.Repository
}

func NewCreateTodoUseCase(repo todo.Repository) *CreateTodo {
	return &CreateTodo{
		repo: repo,
	}
}

func (u *CreateTodo) Execute(ctx context.Context, title string) (*todo.Todo, error) {
	todo := todo.Create(title)

	err := u.repo.Save(ctx, todo)

	if err != nil {
		return nil, err
	}

	return todo, nil
}
