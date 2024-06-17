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

func PostIndex(c *gin.Context) {
	var posts []models.Post
	config.DB.Find(&posts)
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostShow(c *gin.Context) {
	var post models.Post
	config.DB.First(&post, c.Param("id"))
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostUpdate(c *gin.Context) {
	// Get the post id from the url
	id:= c.Param("id")

	// Get the post from the DB
	var body struct {
		Title string
		Body string
	}
	c.Bind(&body)

	// Find the post we are updating
	var post models.Post
	config.DB.First(&post, id)

	// Update it
	config.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})

	// Return it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostDelete(c *gin.Context) {
	// Get the post id from the url
	id:= c.Param("id")

	// Delete the post
	config.DB.Delete(&models.Post{}, id)

	// Respond
	c.Status(200)
}