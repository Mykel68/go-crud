package main

import (
	"go-crud/config"
	"go-crud/models"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectDB()
}

func main() {
	config.DB.AutoMigrate(&models.Post{})
}