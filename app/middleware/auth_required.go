package middleware

import (
	"github.com/iamyoki/go-gin-ddd-clean-architecture/app/appcontext"
	"github.com/iamyoki/go-gin-ddd-clean-architecture/app/apperror"

	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, ok := appcontext.GetActiveUser(c.Request.Context()); !ok {
			c.Error(&apperror.Unauthorized{Msg: "Sign in required"})
			c.Abort()
			return
		}
		c.Next()
	}
}
