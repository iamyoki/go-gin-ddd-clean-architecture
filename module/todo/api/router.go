package api

import "github.com/gin-gonic/gin"

func RegisterRouter(r *gin.RouterGroup, h *todoHandler) {
	g := r.Group("/todos")

	g.POST("", h.HandleCreate)
	g.GET("", h.HandleGetAll)
	g.GET(":id", h.HandleGetById)
	g.POST(":id/complete", h.HandleComplete)
	g.DELETE(":id", h.HandleDelete)
}
