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

	todoHandler := api.NewTodoHandler(createTodoUseCase)

	api.RegisterTodoRouter(m.r, todoHandler)
}
