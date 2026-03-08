package infrastructure

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"
	apperror "todo_api/app/error"
	"todo_api/module/iam/domain/user"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserEntity struct {
	ID             uuid.UUID
	Email          string `gorm:"uniqueIndex;not null"`
	HashedPassword string `gorm:"not null"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (e *UserEntity) toDomain() *user.User {
	return &user.User{
		ID:             e.ID,
		Email:          e.Email,
		HashedPassword: e.HashedPassword,
		CreatedAt:      e.CreatedAt,
		UpdatedAt:      e.UpdatedAt,
	}
}

func fromDomain(user *user.User) UserEntity {
	return UserEntity{
		ID:             user.ID,
		Email:          user.Email,
		HashedPassword: user.HashedPassword,
		CreatedAt:      user.CreatedAt,
		UpdatedAt:      user.UpdatedAt,
	}
}

type GormUserRepository struct {
	DB *gorm.DB
}

// DeleteById implements [user.UserRepositoryInterface].
func (s *GormUserRepository) DeleteById(ctx context.Context, id uuid.UUID) (*user.User, error) {
	panic("unimplemented")
}

// FindAll implements [user.UserRepositoryInterface].
func (s *GormUserRepository) FindAll(ctx context.Context) ([]user.User, error) {
	panic("unimplemented")
}

// FindByEmail implements [user.UserRepositoryInterface].
func (s *GormUserRepository) FindByEmail(ctx context.Context, email string) (*user.User, error) {
	entity, err := gorm.G[UserEntity](s.DB).Where("email = ?", email).First(ctx)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &apperror.NotFound{Msg: "Email not found"}
		}

		return nil, err
	}

	return entity.toDomain(), err
}

// FindById implements [user.UserRepositoryInterface].
func (s *GormUserRepository) FindById(ctx context.Context, id uuid.UUID) (*user.User, error) {
	panic("unimplemented")
}

// Save implements [user.UserRepositoryInterface].
func (s *GormUserRepository) Save(ctx context.Context, user *user.User) error {
	entity := fromDomain(user)
	err := gorm.G[UserEntity](s.DB, clause.OnConflict{UpdateAll: true}).Create(ctx, &entity)

	if err != nil && strings.Contains(err.Error(), "UNIQUE") {
		return &apperror.Conflict{Msg: fmt.Sprintf("The email `%s` already exists", user.Email)}
	}

	return err
}
