package usecase

import (
	"context"
	"todo_api/modules/todo/domain"

	"github.com/google/uuid"
)

type CompleteTodoUseCase struct {
	Repo domain.TodoRepositoryInterface
}

func (usecase *CompleteTodoUseCase) Execute(ctx context.Context, id uuid.UUID) (*domain.Todo, error) {
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
