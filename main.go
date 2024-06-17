package main

import (
	"github.com/gin-gonic/gin"
	"log"
    "github.com/joho/godotenv"
)

func init() {
		  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }


}

func main() {

	  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }


	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}