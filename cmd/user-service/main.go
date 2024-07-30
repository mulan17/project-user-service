package main

import (
	"sign/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	routes.RegisterRoutes(server)
	

	server.Run(":8080") //localhost:8080
}