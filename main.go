package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.GET("/", func(ctx *gin.Context) {

		// ctx.JSON(200, map[string]any{"message": "Todo API is running!", "status": "success"})
		// or
		ctx.JSON(200, gin.H{
			"message": "Todo API is running!",
			"status":  "success",
		},
		)
	})
	router.Run()
}
