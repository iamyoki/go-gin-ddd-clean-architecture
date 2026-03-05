package main

import (
	"todo_api/app/config"
	"todo_api/app/database"
	"todo_api/app/middleware"
	"todo_api/modules/todo"

	"github.com/gin-gonic/gin"
)

type DTO struct {
	Name string `json:"name" binding:"required"`
}

func main() {
	// app setup
	cfg := config.Load()
	db := database.InitSqliteDB(cfg)
	r := gin.Default()
	r.SetTrustedProxies(nil)

	// global middlewares
	r.Use(middleware.ErrorHandler())

	// api root
	api := r.Group("/api")

	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, "ok")
	})

	r.GET("/health3", func(ctx *gin.Context) {
		ctx.JSON(200, "ok")
	})

	// modules setup
	todo.NewModule(db, cfg, api).Register()

	// run server
	r.Run(":" + cfg.AppPort)
}
