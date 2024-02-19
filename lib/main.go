package main

// @title SimpleGoApi API
// @version 1.0
// @description This is a SimpleGoApi server.

// @host localhost:8080
// @BasePath /

import (
	"github.com/gabrielyea/simple-go-api/lib/docs"
	"github.com/gabrielyea/simple-go-api/lib/handlers"
	"net/http"
	"os"

	"context"
	"log"

	"firebase.google.com/go"
	"github.com/gin-gonic/gin"
)

func main() {
	// Set Swagger info
	docs.SwaggerInfo.Title = "Simple Go API"
	docs.SwaggerInfo.Description = "This is a simple API for learning purposes"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http"}

	_, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	r := gin.Default()

	r.POST("/login", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
	})
	r.GET("/users", handlers.GetUsers)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
