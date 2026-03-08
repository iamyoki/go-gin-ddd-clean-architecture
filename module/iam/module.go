package iam

import (
	"todo_api/app/config"
	"todo_api/module/iam/api"
	"todo_api/module/iam/infrastructure"
	"todo_api/module/iam/usecase"

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
	m.db.AutoMigrate(&infrastructure.UserEntity{})

	userRepo := &infrastructure.GormUserRepository{DB: m.db}

	bcryptHasher := &infrastructure.BcryptHasher{}

	signUp := &usecase.SignUp{
		Repo:   userRepo,
		Hasher: bcryptHasher,
	}

	handler := &api.Handler{
		SignUp: signUp,
	}

	api.RegisterRouter(m.r, handler)
}
