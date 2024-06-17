package controllers

import (
	"go-crud/config"
	"go-crud/models"

	"github.com/gin-gonic/gin"
)


func PostCreate (c *gin.Context) {
	// Get data off req body
	var body struct {
		Title string
		Body string
	}

	c.Bind(&body)

	// create post
	post := models.Post{Title: body.Title, Body: body.Body}
	result := config.DB.Create(&post)

if result.Error != nil {
	c.Status(400)
	return
}

// return it
		c.JSON(200, gin.H{
			"post": post,
		})
	}