package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"test-go/initializer"
)

func init() {
	initializer.EnvLoader()
}

func main() {
	fmt.Println("hello")

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}
