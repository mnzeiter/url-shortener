package main

import (
	"log"

	"url-shortener/internal/db"
	"url-shortener/internal/models"
	"url-shortener/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.ConnectDB()

	if err := db.DB.AutoMigrate(&models.URL{}); err != nil {
		log.Fatalf("failed to migrate: %v", err)
	}

	r := gin.Default()

	routes.RegisterRoutes(r)

	log.Println("Server running on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
