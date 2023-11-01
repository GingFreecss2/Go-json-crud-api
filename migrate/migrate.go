package main

import (
	"github.com/GingFreecss2/JSON-CRUD-API-Go/initializers"
	"github.com/GingFreecss2/JSON-CRUD-API-Go/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
