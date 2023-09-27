package main

import (
	"github.com/gin-gonic/gin"
	"github.com/unawaretub86/fiberApiGo/controllers"
	"github.com/unawaretub86/fiberApiGo/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	r := gin.Default()

	r.POST("/post", controllers.PostsCreate)

	r.GET("/posts", controllers.PostIndex)
	r.GET("/post/:id", controllers.PostByID)

	r.PUT("/post/:id", controllers.UpdatePost)

	r.DELETE("/post/:id", controllers.PostDelete)

	r.Run()
}
