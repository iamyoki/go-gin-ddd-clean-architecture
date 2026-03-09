package usecase

import (
	"context"
	"todo_api/module/todo/domain/todo"
)

type CreateTodo struct {
	Repo todo.Repository
}

func (u *CreateTodo) Execute(ctx context.Context, title string) (*todo.Todo, error) {
	todo := todo.Create(title)

	err := u.Repo.Save(ctx, todo)

	if err != nil {
		return nil, err
	}

	return todo, nil
}
