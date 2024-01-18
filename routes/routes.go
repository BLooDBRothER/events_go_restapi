package routes

import (
	"example.com/events/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	authenticated := server.Group("/")
	authenticated.Use(middleware.AuthenticationMiddleware)
	
	// server status route
	server.GET("/status", serverStatus)

	// events routes
	server.GET("/events", GetEvents)
	server.GET("/events/:id", GetEvent)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

	// user routes
	server.POST("/signup", signup)
	server.POST("/login", login)
}
