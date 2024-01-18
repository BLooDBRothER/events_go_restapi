package main

import (
	"github.com/gin-gonic/gin"

	"example.com/events/config"
	"example.com/events/db"
	"example.com/events/routes"
)

func main() {
	config.InitConfig()
	db.InitDB()

	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run()
}
