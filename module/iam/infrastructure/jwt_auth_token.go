package infrastructure

import (
	"context"
	"time"
	"todo_api/app/appcontext"
	"todo_api/app/apperror"
	"todo_api/module/iam/domain/user"
	"todo_api/module/iam/usecase"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTAuthToken struct {
	Secret                []byte
	AccessTokenExpiresIn  time.Duration
	RefreshTokenExpiresIn time.Duration
	UserRepo              user.Repository
}

type Claims struct {
	appcontext.ActiveUser
	jwt.RegisteredClaims
}

func (j *JWTAuthToken) Generate(activeUser appcontext.ActiveUser) (usecase.TokenPair, error) {
	accessClaims := Claims{
		activeUser,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.AccessTokenExpiresIn)),
			ID:        activeUser.ID.String(),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	at, err := accessToken.SignedString(j.Secret)

	if err != nil {
		return usecase.TokenPair{}, err
	}

	refreshClaims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.RefreshTokenExpiresIn)),
		ID:        activeUser.ID.String(),
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	rt, err := refreshToken.SignedString(j.Secret)

	if err != nil {
		return usecase.TokenPair{}, err
	}

	return usecase.TokenPair{
		AccessToken:  at,
		RefreshToken: rt,
	}, nil
}

func (j *JWTAuthToken) Verify(accessToken string) (*appcontext.ActiveUser, error) {
	token, err := jwt.ParseWithClaims(accessToken, &Claims{}, func(t *jwt.Token) (any, error) {
		return []byte(j.Secret), nil
	})

	if err != nil {
		return nil, &apperror.Unauthorized{Msg: "Invalid or expired credentials"}
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return &claims.ActiveUser, nil
	}

	return nil, &apperror.Unauthorized{Msg: "Unknown credentials"}
}

func (j *JWTAuthToken) Refresh(refreshToken string) (usecase.TokenPair, error) {
	token, err := jwt.ParseWithClaims(refreshToken, &jwt.RegisteredClaims{}, func(t *jwt.Token) (any, error) {
		return []byte(j.Secret), nil
	})

	if err != nil {
		return usecase.TokenPair{}, &apperror.Forbidden{Msg: "Invalid refresh token"}
	}

	if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
		id, err := uuid.Parse(claims.ID)
		if err != nil {
			return usecase.TokenPair{}, &apperror.Forbidden{Msg: "Invalid refresh token id"}
		}

		user, err := j.UserRepo.FindById(context.Background(), id)
		if err != nil {
			return usecase.TokenPair{}, &apperror.BadRequest{Msg: "Invalid refresh token user"}
		}

		return j.Generate(user.ToActiveUser())
	}

	return usecase.TokenPair{}, &apperror.Forbidden{Msg: "Invalid refresh token"}
}
