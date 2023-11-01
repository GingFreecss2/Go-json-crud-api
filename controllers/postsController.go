package controllers

import (
	"github.com/GingFreecss2/JSON-CRUD-API-Go/initializers"
	"github.com/GingFreecss2/JSON-CRUD-API-Go/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {

	// Get data off request body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// Create a Post
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	//Return it

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {

	// Get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	// Respond with them
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	// Get the Id off the URL
	id := c.Param("id")

	// Get the Post
	var post models.Post
	initializers.DB.First(&post, id)

	// Respond with the Post
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	// Get the id off the url
	id := c.Param("id")

	// Get the Data off the request body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// Find the Post to update
	var post models.Post
	initializers.DB.First(&post, id)

	// Update the Post
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	// Respond with the Post
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	// Get the id off the URL
	id := c.Param("id")

	// Delete the Post
	initializers.DB.Delete(&models.Post{}, id)

	//Respond
	c.Status(200)
}
