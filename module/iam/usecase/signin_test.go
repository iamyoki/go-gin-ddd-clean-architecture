package usecase_test

import (
	"context"
	"testing"

	"github.com/iamyoki/go-gin-ddd-clean-architecture/app/apperror"
	"github.com/iamyoki/go-gin-ddd-clean-architecture/module/iam/domain/user"
	usermocks "github.com/iamyoki/go-gin-ddd-clean-architecture/module/iam/domain/user/mocks"
	"github.com/iamyoki/go-gin-ddd-clean-architecture/module/iam/usecase"
	usecasemocks "github.com/iamyoki/go-gin-ddd-clean-architecture/module/iam/usecase/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSignIn(t *testing.T) {
	repo := usermocks.NewMockRepository(t)
	hasher := usecasemocks.NewMockHasher(t)
	authToken := usecasemocks.NewMockAuthToken(t)

	signIn := &usecase.SignIn{
		Repo:      repo,
		Hasher:    hasher,
		AuthToken: authToken,
	}

	// valid user
	ctx := context.Background()
	email := "test@example.com"
	pwd := "password123456"
	hashedPwd := "hashedpassword"

	t.Run("sign in ok", func(t *testing.T) {
		// arrange
		mockUser, _ := user.Create(email, hashedPwd)
		mockActiveUser := mockUser.ToActiveUser()
		mockTokenPair := usecase.TokenPair{
			AccessToken:  "someaccesstoken",
			RefreshToken: "somerefreshtoken",
		}

		repo.EXPECT().FindByEmail(ctx, email).Return(mockUser, nil)
		hasher.EXPECT().VerifyPassword(hashedPwd, pwd).Return(true)

		authToken.EXPECT().Generate(mockActiveUser).Return(mockTokenPair, nil)

		// act
		tokenPair, err := signIn.Execute(ctx, email, pwd)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, mockTokenPair, tokenPair)
	})

	t.Run("no user", func(t *testing.T) {
		// arrange
		repo.EXPECT().FindByEmail(ctx, "non-existing@user.com").Return(nil, assert.AnError)

		// act
		_, err := signIn.Execute(ctx, "non-existing@user.com", "somepwd")

		// assert
		assert.Error(t, err)
		assert.IsType(t, &apperror.Unauthorized{}, err)
		assert.Contains(t, err.Error(), "Email does not exsists")
	})

	t.Run("wrong password", func(t *testing.T) {
		// arrange
		mockUser, _ := user.Create(email, hashedPwd)

		repo.EXPECT().FindByEmail(ctx, email).Return(mockUser, nil)
		hasher.EXPECT().VerifyPassword(hashedPwd, "wrong password").Return(false)

		// act
		tokenPair, err := signIn.Execute(ctx, email, "wrong password")

		// assert
		authToken.AssertNotCalled(t, "Generate", mock.Anything)
		var apperr *apperror.Unauthorized
		assert.ErrorAs(t, err, &apperr)
		assert.Contains(t, apperr.Msg, "Incorrect password")
		assert.Equal(t, usecase.TokenPair{}, tokenPair)
	})
}
