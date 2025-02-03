package main

import (
	"test-go/initializer"
	"test-go/models"
)

func init() {
	initializer.LoadEnv()
	initializer.ConnectToDB()
}

func main() {
	initializer.DB.AutoMigrate(&models.User{})
}
