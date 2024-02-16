package main

import (
	"fmt"

	"example.com/rest-api/db"
	"example.com/rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Welcome to Go Rest-API Project")

	db.InitDB()

	server := gin.Default()

	routes.RegisterRouts(server)

	server.Run(":8080")
}
