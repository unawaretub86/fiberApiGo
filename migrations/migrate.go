package main

import (
	"github.com/unawaretub86/fiberApiGo/initializers"
	"github.com/unawaretub86/fiberApiGo/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
