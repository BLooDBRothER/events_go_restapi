package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"example.com/events/models"
)

func main() {
	server := gin.Default()

	server.GET("/status", serverStatus)

	server.GET("/events", )

	server.Run();
}

func serverStatus(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Server is running..."})
}

func getEvents(context *gin.Context) {
	context.JSON(http.StatusOK, models.GetEvents())
}
