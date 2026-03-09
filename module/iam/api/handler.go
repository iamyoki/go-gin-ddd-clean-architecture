package api

import (
	"net/http"
	"time"

	"github.com/iamyoki/go-gin-ddd-clean-architecture/module/iam/constant"
	"github.com/iamyoki/go-gin-ddd-clean-architecture/module/iam/domain/user"
	"github.com/iamyoki/go-gin-ddd-clean-architecture/module/iam/usecase"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handler struct {
	SignUp *usecase.SignUp
	SignIn *usecase.SignIn
}

type SignUpDTO struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserDTO struct {
	Id        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SignInDTO struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type TokenPairDTO struct {
	AccessToken  string `json:"token_access"`
	RefreshToken string `json:"refresh_token"`
}

func fromDomain(user *user.User) UserDTO {
	return UserDTO{
		Id:        user.ID,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func (h *Handler) HandleSignUp(c *gin.Context) {
	var dto SignUpDTO

	if err := c.ShouldBind(&dto); err != nil {
		c.Error(err)
		return
	}

	user, err := h.SignUp.Execute(c.Request.Context(), dto.Email, dto.Password)

	if err != nil {
		c.Error(err)
		return
	}

	userDTO := fromDomain(user)

	c.JSON(http.StatusCreated, userDTO)
}

func (h *Handler) HandleSignIn(c *gin.Context) {
	var dto SignInDTO

	if err := c.ShouldBind(&dto); err != nil {
		c.Error(err)
		return
	}

	tokenPair, err := h.SignIn.Execute(c.Request.Context(), dto.Email, dto.Password)

	if err != nil {
		c.Error(err)
		return
	}

	tokenPairDTO := TokenPairDTO{
		AccessToken:  tokenPair.AccessToken,
		RefreshToken: tokenPair.RefreshToken,
	}

	c.SetCookie("access_token", tokenPair.AccessToken, int(constant.AccessTokenExpiresIn.Seconds()), "/", "", false, true)
	c.SetCookie("refresh_token", tokenPair.RefreshToken, int(constant.RefreshTokenExpiresIn.Seconds()), "/", "", false, true)
	c.JSON(http.StatusOK, tokenPairDTO)
}
