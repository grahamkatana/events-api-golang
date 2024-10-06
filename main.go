package main

import (
	"grahamkatana/api/events/db"
	"grahamkatana/api/events/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	server.GET("/", healthCheck)
	routes.BootstrapApiRoutes(server)
	server.Run(":8080")

}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "The server is running",
	})
}
