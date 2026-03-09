package usecase

import (
	"todo_api/app/appcontext"
)

type Hasher interface {
	HashPassowrd(password string) (string, error)
	VerifyPassword(hashedPassword string, password string) bool
}

type TokenPair struct {
	AccessToken  string
	RefreshToken string
}

type AuthToken interface {
	Generate(activeUser appcontext.ActiveUser) (TokenPair, error)
	Verify(accessToken string) (*appcontext.ActiveUser, error)
	Refresh(refreshToken string) (TokenPair, error)
}
