package middleware

import (
	apperror "todo_api/app/error"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			// capture last err
			err := c.Errors.Last().Err

			// turn into response
			status, body := apperror.IntoResponse(err)

			// send
			c.JSON(status, body)
		}
	}
}
