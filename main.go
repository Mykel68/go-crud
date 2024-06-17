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
	r.GET("/posts", controllers.PostIndex)
	r.GET("/post/:id", controllers.PostShow)
	r.PUT("/post/:id", controllers.PostUpdate)
	r.DELETE("/post/:id", controllers.PostDelete)
	r.Run() 
}