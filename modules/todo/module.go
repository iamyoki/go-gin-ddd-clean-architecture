package todo

import (
	"todo_api/app/config"
	"todo_api/modules/todo/api"
	"todo_api/modules/todo/infrastructure"
	"todo_api/modules/todo/usecase"

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

	todoRepo := infrastructure.NewSqliteTodoRepository(m.db)

	createTodoUseCase := usecase.NewCreateTodoUseCase(todoRepo)

	getAllTodosUseCase := usecase.NewGetAllTodosUseCase(todoRepo)

	getTodoByIdUseCase := &usecase.GetTodoByIdUseCase{Repo: todoRepo}

	completeTodoUseCase := &usecase.CompleteTodoUseCase{Repo: todoRepo}

	todoHandler := api.NewTodoHandler(createTodoUseCase, getAllTodosUseCase, getTodoByIdUseCase, completeTodoUseCase)

	api.RegisterTodoRouter(m.r, todoHandler)
}
