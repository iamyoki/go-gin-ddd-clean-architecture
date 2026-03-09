package iam

import (
	"todo_api/app/config"
	"todo_api/module/iam/api"
	"todo_api/module/iam/constant"
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

	// infra
	userRepo := &infrastructure.GormUserRepository{DB: m.db}
	bcryptHasher := &infrastructure.BcryptHasher{}
	jwtAuthToken := &infrastructure.JWTAuthToken{
		Secret:                []byte(m.config.JWTSecret),
		AccessTokenExpiresIn:  constant.AccessTokenExpiresIn,
		RefreshTokenExpiresIn: constant.RefreshTokenExpiresIn,
		UserRepo:              userRepo,
	}

	// usecase
	signUp := &usecase.SignUp{
		Repo:   userRepo,
		Hasher: bcryptHasher,
	}
	signIn := &usecase.SignIn{
		Repo:      userRepo,
		Hasher:    bcryptHasher,
		AuthToken: jwtAuthToken,
	}

	// api
	handler := &api.Handler{
		SignUp: signUp,
		SignIn: signIn,
	}

	api.RegisterRouter(m.r, handler)
}
