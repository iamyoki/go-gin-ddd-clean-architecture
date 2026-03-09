package usecase

import (
	"context"
	"log/slog"

	"github.com/iamyoki/go-gin-ddd-clean-architecture/app/apperror"
	"github.com/iamyoki/go-gin-ddd-clean-architecture/module/iam/domain/user"
)

type SignIn struct {
	Repo      user.Repository
	Hasher    Hasher
	AuthToken AuthToken
}

func (u *SignIn) Execute(ctx context.Context, email, password string) (TokenPair, error) {
	// check user exists
	user, err := u.Repo.FindByEmail(ctx, email)
	if err != nil {
		return TokenPair{}, &apperror.Unauthorized{Msg: "Email does not exsists"}
	}

	// verify password
	ok := u.Hasher.VerifyPassword(user.HashedPassword, password)
	if !ok {
		return TokenPair{}, &apperror.Unauthorized{Msg: "Incorrect password"}
	}

	// generate token pair
	tokenPair, err := u.AuthToken.Generate(user.ToActiveUser())
	if err != nil {
		return TokenPair{}, err
	}

	slog.Info("User logged in", "id", user.ID, "access_token", tokenPair.AccessToken)

	return tokenPair, nil
}
