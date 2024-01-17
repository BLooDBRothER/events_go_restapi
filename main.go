package main

import (
	"github.com/gin-gonic/gin"

	"example.com/events/db"
	"example.com/events/routes"
)

func main() {
	db.InitDB()

	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run()
}
