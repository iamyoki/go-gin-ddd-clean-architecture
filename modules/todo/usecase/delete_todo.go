package usecase

import (
	"context"
	"todo_api/modules/todo/domain"

	"github.com/google/uuid"
)

type DeleteTodoUseCase struct {
	Repo domain.TodoRepositoryInterface
}

func (usecase *DeleteTodoUseCase) Execute(ctx context.Context, id uuid.UUID) (*domain.Todo, error) {
	return usecase.Repo.DeleteById(ctx, id)
}
