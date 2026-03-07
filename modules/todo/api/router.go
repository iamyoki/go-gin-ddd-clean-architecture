package api

import "github.com/gin-gonic/gin"

func RegisterTodoRouter(r *gin.RouterGroup, h *todoHandler) {
	g := r.Group("/todos")

	g.POST("", h.Create)
	g.GET("", h.GetAll)
	g.GET(":id", h.GetById)
	g.POST(":id/complete", h.Complete)
}
