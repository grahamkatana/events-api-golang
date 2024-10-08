package main

import (
	"grahamkatana/api/events/db"
	"grahamkatana/api/events/routes"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		panic("Failed to load .env file")
	}
	db.InitDB()
	server := gin.Default()
	server.GET("/", healthCheck)
	routes.BootstrapAuthApiRoutes(server)
	routes.BootstrapApiRoutes(server)
	server.Run(":3050")

}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "The server is running",
	})
}
