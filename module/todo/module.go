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
	db      *gorm.DB
	config  *config.Config
	handler *api.Handler
}

func NewModule(db *gorm.DB, config *config.Config) *module {
	// infra
	todoRepo := infrastructure.NewGormTodoRepository(db)

	// usecase
	createTodo := &usecase.CreateTodo{Repo: todoRepo}
	getAllTodos := &usecase.GetAllTodos{Repo: todoRepo}
	getTodo := &usecase.GetTodo{Repo: todoRepo}
	completeTodo := &usecase.CompleteTodo{Repo: todoRepo}
	deleteTodo := &usecase.DeleteTodo{Repo: todoRepo}

	// api
	handler := &api.Handler{
		CreateTodo:   createTodo,
		GetAllTodos:  getAllTodos,
		GetTodo:      getTodo,
		CompleteTodo: completeTodo,
		DeleteTodo:   deleteTodo,
	}

	return &module{
		db:      db,
		config:  config,
		handler: handler,
	}
}

func (m *module) Init(r *gin.RouterGroup) {
	m.db.AutoMigrate(&infrastructure.TodoEntity{})
	api.RegisterRouter(r, m.handler)
}
