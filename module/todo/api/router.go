package api

import (
	"github.com/iamyoki/go-gin-ddd-clean-architecture/app/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.RouterGroup, h *Handler) {
	g := r.Group("/todos")

	g.POST("", h.HandleCreate)
	g.GET("", h.HandleGetAll)
	g.GET(":id", h.HandleGetById)
	g.POST(":id/complete", h.HandleComplete)
	g.DELETE(":id", middleware.AuthRequired(), h.HandleDelete)
}
