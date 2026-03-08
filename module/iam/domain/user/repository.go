package user

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	Save(ctx context.Context, user *User) error
	FindAll(ctx context.Context) ([]User, error)
	FindById(ctx context.Context, id uuid.UUID) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	DeleteById(ctx context.Context, id uuid.UUID) (*User, error)
}
