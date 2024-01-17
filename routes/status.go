package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func serverStatus(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Server is running..."})
}
