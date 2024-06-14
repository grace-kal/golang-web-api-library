package main

import (
	"WebApiLibrary/database"
	"WebApiLibrary/models"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	database.InitDb()
	router := gin.Default()

	// Routes
	router.GET("/books", models.GetAllBooks)
	router.GET("/books/:id", models.GetBookByID)
	router.POST("/books", models.CreateBook)
	router.PUT("/books/:id", models.UpdateBook)
	router.DELETE("/books/:id", models.DeleteBook)

	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
