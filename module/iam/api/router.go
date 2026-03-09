package api

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.RouterGroup, h *Handler) {
	g := r.Group("auth")

	g.POST("sign-up", h.HandleSignUp)
	g.POST("sign-in", h.HandleSignIn)
}
