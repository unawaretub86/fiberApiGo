package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/unawaretub86/fiberApiGo/initializers"
	"github.com/unawaretub86/fiberApiGo/models"
)

func PostsCreate(c *gin.Context) {
	//get data from req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	//create post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)
	if result.Error != nil {
		c.Status(400)
		return
	}

	//return it
	c.JSON(201, gin.H{
		"post": post,
	})
}

func PostIndex(c *gin.Context) {
	// Get post
	var posts []models.Post
	initializers.DB.Find(&posts)

	// returning it
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostByID(c *gin.Context) {
	// Get id off url
	id := c.Param("id")

	// Get post
	var post models.Post
	initializers.DB.First(&post, id)

	//returning it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func UpdatePost(c *gin.Context) {
	// Get id from url
	id := c.Param("id")

	// Get data from req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// Find the post to update
	var post models.Post
	initializers.DB.First(&post, id)

	// Update the post
	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})

	// Return
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostDelete(c *gin.Context) {
	// Get id from url
	id := c.Param("id")

	// Update the post
	initializers.DB.Delete(&models.Post{}, id)

	// Return
	c.JSON(200, gin.H{
		"post": "Deleted successfully",
	})
}
