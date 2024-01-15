package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	err := server.Run(":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		panic(err)
	}
}
