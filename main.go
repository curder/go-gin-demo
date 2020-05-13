package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	var (
		r *gin.Engine
	)
	r = gin.Default()

	r.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	panic(r.Run())
}
