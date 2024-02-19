package main

// @title SimpleGoApi API
// @version 1.0
// @description This is a SimpleGoApi server.

// @host localhost:8080
// @BasePath /

import (
	handlers "github/gabrielyea/simple-go-api/api"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger/example/basic/docs"
)

func main() {
	// Set Swagger info
	docs.SwaggerInfo_swagger.Title = "Simple Go API"
	docs.SwaggerInfo_swagger.Description = "This is a simple API for learning purposes"
	docs.SwaggerInfo_swagger.Version = "1.0"
	docs.SwaggerInfo_swagger.BasePath = "/"
	docs.SwaggerInfo_swagger.Schemes = []string{"http"}

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
