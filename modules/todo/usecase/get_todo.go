package usecase

import (
	"context"
	"todo_api/modules/todo/domain"

	"github.com/google/uuid"
)

type GetTodoUseCase struct {
	Repo domain.TodoRepositoryInterface
}

func (usecase *GetTodoUseCase) Execute(ctx context.Context, id uuid.UUID) (*domain.Todo, error) {
	return usecase.Repo.FindById(ctx, id)
}
