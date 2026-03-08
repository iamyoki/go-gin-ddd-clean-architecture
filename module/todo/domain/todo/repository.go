package todo

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	Save(ctx context.Context, todo *Todo) error
	FindAll(ctx context.Context) ([]Todo, error)
	FindById(ctx context.Context, id uuid.UUID) (*Todo, error)
	DeleteById(ctx context.Context, id uuid.UUID) (*Todo, error)
}
