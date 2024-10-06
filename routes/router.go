package routes

import (
	"grahamkatana/api/events/controllers"

	"github.com/gin-gonic/gin"
)

func BootstrapApiRoutes(server *gin.Engine) {
	api := server.Group("/api")
	{
		api.GET("/events", controllers.GetEvents)
		api.GET("/events/:id", controllers.GetEvent)
		api.POST("/events", controllers.CreateEvent)
		api.PATCH("/events/:id", controllers.UpdateEvent)
		api.DELETE("/events/:id", controllers.DeleteEvent)
	}
}
