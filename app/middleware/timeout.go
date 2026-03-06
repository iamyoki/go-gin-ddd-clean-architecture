package middleware

import (
	"time"

	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
)

func Timeout(t time.Duration) gin.HandlerFunc {
	return timeout.New(timeout.WithTimeout(t))
}
