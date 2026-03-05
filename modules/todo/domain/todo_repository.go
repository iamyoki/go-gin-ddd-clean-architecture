package domain

import "context"

type TodoRepositoryInterface interface {
	Save(ctx context.Context, todo *Todo) error
}
