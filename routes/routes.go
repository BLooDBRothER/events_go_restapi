package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	// server status route
	server.GET("/status", serverStatus)

	// events routes
	server.GET("/events", GetEvents)
	server.GET("/events/:id", GetEvent)
	server.POST("/events", createEvent)
	server.PUT("/events/:id", updateEvent)
	server.DELETE("/events/:id", deleteEvent)

	// user routes
	server.POST("/signup", signup)
	server.POST("/login", login)
}
