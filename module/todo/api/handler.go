package api

import (
	"net/http"
	"time"
	"todo_api/app/apperror"
	"todo_api/module/todo/usecase"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type todoHandler struct {
	createTodoUseCase   *usecase.CreateTodo
	getAllTodosUseCase  *usecase.GetAllTodos
	getTodoByIdUseCase  *usecase.GetTodo
	completeTodoUseCase *usecase.CompleteTodo
	deleteTodoUseCase   *usecase.DeleteTodo
}

func NewTodoHandler(
	createTodoUseCase *usecase.CreateTodo,
	getAllTodosUseCase *usecase.GetAllTodos,
	getTodoByIdUseCase *usecase.GetTodo,
	completeTodoUseCase *usecase.CompleteTodo,
	deleteTodoUseCase *usecase.DeleteTodo,
) *todoHandler {
	return &todoHandler{
		createTodoUseCase:   createTodoUseCase,
		getAllTodosUseCase:  getAllTodosUseCase,
		getTodoByIdUseCase:  getTodoByIdUseCase,
		completeTodoUseCase: completeTodoUseCase,
		deleteTodoUseCase:   deleteTodoUseCase,
	}
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

func (h *todoHandler) HandleCreate(c *gin.Context) {
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
		Id:          todo.ID,
		Title:       todo.Title,
		Completed:   todo.Completed,
		CreatedAt:   todo.CreatedAt,
		UpdatedAt:   todo.UpdatedAt,
		CompletedAt: todo.CompletedAt,
	})

}

func (h *todoHandler) HandleGetAll(c *gin.Context) {
	todos, err := h.getAllTodosUseCase.Execute(c.Request.Context())

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

func (h *todoHandler) HandleGetById(c *gin.Context) {
	idStr := c.Param("id")

	id, err := uuid.Parse(idStr)

	if err != nil {
		c.Error(&apperror.BadRequest{Msg: "Invalid ID"})
		return
	}

	todo, err := h.getTodoByIdUseCase.Execute(c.Request.Context(), id)

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

func (h *todoHandler) HandleComplete(c *gin.Context) {
	idStr := c.Param("id")

	id, err := uuid.Parse(idStr)

	if err != nil {
		c.Error(&apperror.BadRequest{Msg: "Invalid ID"})
		return
	}

	todo, err := h.completeTodoUseCase.Execute(c.Request.Context(), id)

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

func (h *todoHandler) HandleDelete(c *gin.Context) {
	idStr := c.Param("id")

	id, err := uuid.Parse(idStr)

	if err != nil {
		c.Error(&apperror.BadRequest{Msg: "Invalid ID"})
		return
	}

	todo, err := h.deleteTodoUseCase.Execute(c.Request.Context(), id)

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
