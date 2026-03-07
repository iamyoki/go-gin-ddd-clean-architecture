package api

import (
	"net/http"
	"time"
	"todo_api/modules/todo/usecase"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type todoHandler struct {
	createTodoUseCase  *usecase.CreateTodoUseCase
	getAllTodosUseCase *usecase.GetAllTodosUseCase
}

func NewTodoHandler(createTodoUseCase *usecase.CreateTodoUseCase, getAllTodosUseCase *usecase.GetAllTodosUseCase) *todoHandler {
	return &todoHandler{
		createTodoUseCase:  createTodoUseCase,
		getAllTodosUseCase: getAllTodosUseCase,
	}
}

type CreateDTO struct {
	Title string `json:"title" binding:"required"`
}

type ResponseDTO struct {
	Id        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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

	c.JSON(http.StatusCreated, ResponseDTO{
		Id:        todo.ID,
		Title:     todo.Title,
		Completed: todo.Completed,
		CreatedAt: todo.CreatedAt,
		UpdatedAt: todo.UpdatedAt,
	})
}

func (h *todoHandler) GetAll(c *gin.Context) {
	todos, err := h.getAllTodosUseCase.Execute(c)

	if err != nil {
		c.Error(err)
		return
	}

	data := make([]ResponseDTO, 0, len(todos))

	for _, todo := range todos {
		data = append(data, ResponseDTO{
			Id:        todo.ID,
			Title:     todo.Title,
			Completed: todo.Completed,
			CreatedAt: todo.CreatedAt,
			UpdatedAt: todo.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, data)
}
