package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/status", serverStatus)

	server.Run();
}

func serverStatus(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Server is running..."})
}
