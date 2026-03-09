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
	db                 *gorm.DB
	config             *config.Config
	r                  *gin.RouterGroup
	handler            *api.Handler
	identityMiddleware gin.HandlerFunc
}

func NewModule(db *gorm.DB, config *config.Config) *module {
	// infra
	userRepo := &infrastructure.GormUserRepository{DB: db}
	bcryptHasher := &infrastructure.BcryptHasher{}
	jwtAuthToken := &infrastructure.JWTAuthToken{
		Secret:                []byte(config.JWTSecret),
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

	identityMiddleware := api.IdentityMiddleware(jwtAuthToken)

	return &module{
		db:                 db,
		config:             config,
		handler:            handler,
		identityMiddleware: identityMiddleware,
	}
}

func (m *module) Init(r *gin.RouterGroup) {
	m.db.AutoMigrate(&infrastructure.UserEntity{})
	r.Use(m.identityMiddleware)
	api.RegisterRouter(r, m.handler)
}

func (m *module) IdentityMiddleware() gin.HandlerFunc {
	return m.identityMiddleware
}
