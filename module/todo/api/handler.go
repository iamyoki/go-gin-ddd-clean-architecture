package api

import (
	"net/http"
	"time"

	"github.com/iamyoki/go-gin-ddd-clean-architecture/app/apperror"
	"github.com/iamyoki/go-gin-ddd-clean-architecture/module/todo/usecase"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handler struct {
	CreateTodo   *usecase.CreateTodo
	GetAllTodos  *usecase.GetAllTodos
	GetTodo      *usecase.GetTodo
	CompleteTodo *usecase.CompleteTodo
	DeleteTodo   *usecase.DeleteTodo
}

type CreateDTO struct {
	Title string `json:"title" binding:"required"`
}

type ResponseDTO struct {
	Id          uuid.UUID  `json:"id"`
	Title       string     `json:"title"`
	Completed   bool       `json:"completed"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	CompletedAt *time.Time `json:"completed_at"`
}

func (h *Handler) HandleCreate(c *gin.Context) {
	var dto CreateDTO

	if err := c.ShouldBind(&dto); err != nil {
		c.Error(err)
		return
	}

	todo, err := h.CreateTodo.Execute(c.Request.Context(), dto.Title)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, ResponseDTO{
		Id:          todo.ID,
		Title:       todo.Title,
		Completed:   todo.Completed,
		CreatedAt:   todo.CreatedAt,
		UpdatedAt:   todo.UpdatedAt,
		CompletedAt: todo.CompletedAt,
	})

}

func (h *Handler) HandleGetAll(c *gin.Context) {
	todos, err := h.GetAllTodos.Execute(c.Request.Context())

	if err != nil {
		c.Error(err)
		return
	}

	data := make([]ResponseDTO, 0, len(todos))

	for _, todo := range todos {
		data = append(data, ResponseDTO{
			Id:          todo.ID,
			Title:       todo.Title,
			Completed:   todo.Completed,
			CreatedAt:   todo.CreatedAt,
			UpdatedAt:   todo.UpdatedAt,
			CompletedAt: todo.CompletedAt,
		})
	}

	c.JSON(http.StatusOK, data)
}

func (h *Handler) HandleGetById(c *gin.Context) {
	idStr := c.Param("id")

	id, err := uuid.Parse(idStr)

	if err != nil {
		c.Error(&apperror.BadRequest{Msg: "Invalid ID"})
		return
	}

	todo, err := h.GetTodo.Execute(c.Request.Context(), id)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, ResponseDTO{
		Id:          todo.ID,
		Title:       todo.Title,
		Completed:   todo.Completed,
		CreatedAt:   todo.CreatedAt,
		UpdatedAt:   todo.UpdatedAt,
		CompletedAt: todo.CompletedAt,
	})
}

func (h *Handler) HandleComplete(c *gin.Context) {
	idStr := c.Param("id")

	id, err := uuid.Parse(idStr)

	if err != nil {
		c.Error(&apperror.BadRequest{Msg: "Invalid ID"})
		return
	}

	todo, err := h.CompleteTodo.Execute(c.Request.Context(), id)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, ResponseDTO{
		Id:          todo.ID,
		Title:       todo.Title,
		Completed:   todo.Completed,
		CreatedAt:   todo.CreatedAt,
		UpdatedAt:   todo.UpdatedAt,
		CompletedAt: todo.CompletedAt,
	})
}

func (h *Handler) HandleDelete(c *gin.Context) {
	idStr := c.Param("id")

	id, err := uuid.Parse(idStr)

	if err != nil {
		c.Error(&apperror.BadRequest{Msg: "Invalid ID"})
		return
	}

	todo, err := h.DeleteTodo.Execute(c.Request.Context(), id)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, ResponseDTO{
		Id:          todo.ID,
		Title:       todo.Title,
		Completed:   todo.Completed,
		CreatedAt:   todo.CreatedAt,
		UpdatedAt:   todo.UpdatedAt,
		CompletedAt: todo.CompletedAt,
	})
}
