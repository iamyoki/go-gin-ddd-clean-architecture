package usecase

import (
	"context"
	"todo_api/modules/todo/domain"
)

type GetAllTodosUseCase struct {
	repo domain.TodoRepositoryInterface
}

func NewGetAllTodosUseCase(repo domain.TodoRepositoryInterface) *GetAllTodosUseCase {
	return &GetAllTodosUseCase{repo: repo}
}

func (usecase *GetAllTodosUseCase) Execute(ctx context.Context) ([]domain.Todo, error) {
	todos, err := usecase.repo.FindAll(ctx)

	if err != nil {
		return nil, err
	}

	return todos, nil
}
