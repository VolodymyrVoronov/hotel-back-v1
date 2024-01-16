package main

import (
	"fmt"
	"hotel-back-v1/db"
	"hotel-back-v1/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	server := gin.Default()

	routes.RegisterRoutes(server)

	err := server.Run(":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		panic(err)
	}
}
