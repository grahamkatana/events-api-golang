package routes

import (
	"grahamkatana/api/events/controllers"
	"grahamkatana/api/events/middleware"

	"github.com/gin-gonic/gin"
)

func BootstrapApiRoutes(server *gin.Engine) {
	api := server.Group("/api/v1")
	authenticated := api.Group("/")
	authenticated.Use(middleware.CheckIsTokenValid)

	api.GET("/events", controllers.GetEvents)
	api.GET("/events/:id", controllers.GetEvent)
	authenticated.POST("/events", controllers.CreateEvent)
	authenticated.PATCH("/events/:id", controllers.UpdateEvent)
	authenticated.DELETE("/events/:id", controllers.DeleteEvent)
	authenticated.POST("/events/register/:eventId", controllers.BookEvent)
	authenticated.DELETE("/events/register/:id", controllers.CancelEvent)

}

func BootstrapAuthApiRoutes(server *gin.Engine) {
	api := server.Group("/api/v1/auth")
	{
		api.POST("/login", controllers.Login)
		api.POST("/register", controllers.Register)
	}
}
