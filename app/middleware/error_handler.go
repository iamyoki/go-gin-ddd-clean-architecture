package middleware

import (
	"fmt"
	"todo_api/app/apperror"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		fmt.Println(c.Errors)

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
