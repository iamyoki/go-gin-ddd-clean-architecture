package usecase

import (
	"context"
	"todo_api/modules/todo/domain"

	"github.com/google/uuid"
)

type GetTodoByIdUseCase struct {
	Repo domain.TodoRepositoryInterface
}

func (usecase *GetTodoByIdUseCase) Execute(ctx context.Context, id uuid.UUID) (*domain.Todo, error) {
	return usecase.Repo.FindById(ctx, id)
}
