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

	// global middlewares
	r.Use(gin.Recovery())
	r.Use(sloggin.New(logger))
	r.Use(middleware.ErrorHandler())
	r.Use(middleware.LimitMax(50)) // 50MB

	// api root
	api := r.Group("/api")

	api.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, "ok")
	})

	// modules setup
	todo.NewModule(db, cfg, api).Register()
	iam.NewModule(db, cfg, api).Register()

	// run server
	r.Run(":" + cfg.AppPort)
}
