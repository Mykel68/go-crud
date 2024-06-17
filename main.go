package main

import (
	"go-crud/config"
	"go-crud/controllers"

	"github.com/gin-gonic/gin"
)

func init() {
config.LoadEnvVariables()
config.ConnectDB()

}

func main() {
	r := gin.Default()
	r.POST("/post", controllers.PostCreate)
	r.Run() 
}