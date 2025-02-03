package main

import (
	"github.com/gin-gonic/gin"

	"test-go/initializer"
)

func init() {
	initializer.LoadEnv()
	initializer.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}
