package api

import (
	"strings"
	"todo_api/module/iam/usecase"

	"github.com/gin-gonic/gin"
)

func IdentityMiddleware(authToken usecase.AuthToken) gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string

		// get token from Authorization header
		authHeader := c.GetHeader("Authorization")
		const Prefix string = "Bearer "
		if authHeader != "" && strings.HasPrefix(authHeader, Prefix) {
			token = strings.TrimPrefix(authHeader, Prefix)
		}

		//  or from cookie
		if token == "" {
			if v, err := c.Cookie("access_token"); err == nil {
				token = v
			}
		}

		// pass ActiveUser into context
		if token != "" {
			if activeUser, err := authToken.Verify(token); err == nil {
				c.Request = c.Request.WithContext(activeUser.IntoContext(c.Request.Context()))
			}
		}

		c.Next()
	}
}
