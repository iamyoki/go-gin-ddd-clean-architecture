package usecase

import (
	"context"
	"todo_api/app/apperror"
	"todo_api/module/iam/domain/user"
)

type SignUp struct {
	Repo   user.Repository
	Hasher Hasher
}

func (u *SignUp) Execute(ctx context.Context, email, password string) (*user.User, error) {
	if err := user.ValidatePassword(password); err != nil {
		return nil, err
	}

	existingUser, err := u.Repo.FindByEmail(ctx, email)
	if _, ok := err.(*apperror.NotFound); !ok {
		return nil, err
	}
	if existingUser != nil {
		return nil, &apperror.Conflict{Msg: "Email already registered"}
	}

	hashedPassword, err := u.Hasher.HashPassowrd(password)
	if err != nil {
		return nil, err
	}

	user, err := user.Create(email, hashedPassword)
	if err != nil {
		return nil, err
	}

	if err := u.Repo.Save(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}
