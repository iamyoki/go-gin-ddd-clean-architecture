package todo

import (
	"todo_api/app/config"
	"todo_api/module/todo/api"
	"todo_api/module/todo/infrastructure"
	"todo_api/module/todo/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type module struct {
	db     *gorm.DB
	config *config.Config
	r      *gin.RouterGroup
}

func NewModule(db *gorm.DB, config *config.Config, r *gin.RouterGroup) *module {
	return &module{
		db:     db,
		config: config,
		r:      r,
	}
}

func (m *module) Register() {
	m.db.AutoMigrate(&infrastructure.TodoEntity{})

	todoRepo := infrastructure.NewGormTodoRepository(m.db)

	createTodoUseCase := usecase.NewCreateTodoUseCase(todoRepo)
	getAllTodosUseCase := usecase.NewGetAllTodosUseCase(todoRepo)
	getTodoUseCase := &usecase.GetTodo{Repo: todoRepo}
	completeTodoUseCase := &usecase.CompleteTodo{Repo: todoRepo}
	deleteTodoUseCase := &usecase.DeleteTodo{Repo: todoRepo}

	todoHandler := api.NewTodoHandler(createTodoUseCase, getAllTodosUseCase, getTodoUseCase, completeTodoUseCase, deleteTodoUseCase)

	api.RegisterRouter(m.r, todoHandler)
}
