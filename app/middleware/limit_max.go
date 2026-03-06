package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LimitMax(maxMB float64) gin.HandlerFunc {
	const MB = 1 << 20
	limitSize := int64(maxMB * float64(MB))
	return func(c *gin.Context) {
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, limitSize)
		c.Next()
	}
}
