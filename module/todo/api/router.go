package api

import (
	"todo_api/app/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.RouterGroup, h *Handler) {
	g := r.Group("/todos", middleware.AuthRequired())

	g.POST("", h.HandleCreate)
	g.GET("", h.HandleGetAll)
	g.GET(":id", h.HandleGetById)
	g.POST(":id/complete", h.HandleComplete)
	g.DELETE(":id", h.HandleDelete)
}
