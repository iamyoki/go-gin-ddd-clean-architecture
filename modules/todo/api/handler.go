package api

import (
	"net/http"
	"todo_api/modules/todo/usecase"

	"github.com/gin-gonic/gin"
)

type todoHandler struct {
	createTodoUseCase *usecase.CreateTodoUseCase
}

func NewTodoHandler(createTodoUseCase *usecase.CreateTodoUseCase) *todoHandler {
	return &todoHandler{
		createTodoUseCase: createTodoUseCase,
	}
}

type CreateDTO struct {
	Title string `json:"title" binding:"required"`
}

func (h *todoHandler) Create(c *gin.Context) {
	var dto CreateDTO

	if err := c.ShouldBind(&dto); err != nil {
		c.Error(err)
		return
	}

	todo, err := h.createTodoUseCase.Execute(c.Request.Context(), dto.Title)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         todo.ID,
		"title":      todo.Title,
		"completed":  todo.Completed,
		"created_at": todo.CreatedAt,
		"updated_at": todo.UpdatedAt,
	})
}
