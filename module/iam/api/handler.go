package api

import (
	"net/http"
	"time"
	"todo_api/module/iam/domain/user"
	"todo_api/module/iam/usecase"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handler struct {
	SignUp *usecase.SignUp
}

type SignUpDTO struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ResponseDTO struct {
	Id        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func fromDomain(user *user.User) ResponseDTO {
	return ResponseDTO{
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

	createdUser := fromDomain(user)

	c.JSON(http.StatusCreated, createdUser)
}
