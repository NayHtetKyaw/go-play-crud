package main

import (
	"test-go/initializer"
	"test-go/models"
)

func init() {
	initializer.EnvLoader()
	initializer.ConnectDB()
}

func main() {
	initializer.DB.AutoMigrate(models.User{})
}
