package main

import (
	"todo_api/app/config"
	"todo_api/app/database"
	"todo_api/app/logger"
	"todo_api/app/middleware"
	"todo_api/module/iam"
	"todo_api/module/todo"

	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
)

type DTO struct {
	Name string `json:"name" binding:"required"`
}

func main() {
	// app setup
	cfg := config.Load()
	db := database.InitSqliteDB(cfg)
	logger := logger.InitLogger()
	r := gin.New()
	r.SetTrustedProxies(nil)

	// module
	iam := iam.NewModule(db, cfg)
	todo := todo.NewModule(db, cfg)

	// global middlewares
	r.Use(gin.Recovery())
	r.Use(sloggin.New(logger))
	r.Use(middleware.ErrorHandler())
	r.Use(middleware.LimitMax(50)) // 50MB
	r.Use(iam.IdentityMiddleware())

	// api root
	api := r.Group("/api")

	// module init
	iam.Init(api)
	todo.Init(api)

	// run server
	r.Run(":" + cfg.AppPort)
}
