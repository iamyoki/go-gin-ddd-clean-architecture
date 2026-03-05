package usecase

import (
	"context"
	"todo_api/modules/todo/domain"
)

type CreateTodoUseCase struct {
	repo domain.TodoRepositoryInterface
}

func NewCreateTodoUseCase(repo domain.TodoRepositoryInterface) *CreateTodoUseCase {
	return &CreateTodoUseCase{
		repo: repo,
	}
}

func (u *CreateTodoUseCase) Execute(ctx context.Context, title string) (*domain.Todo, error) {
	todo := domain.Create(title)

	err := u.repo.Save(ctx, todo)

	if err != nil {
		return nil, err
	}

	return todo, nil
}
