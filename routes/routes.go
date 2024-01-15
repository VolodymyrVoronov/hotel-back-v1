package routes

import (
	"hotel-back-v1/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.POST("/signup", Signup)
	server.POST("/login", Login)

	authenticated := server.Group("/admin")
	authenticated.Use(middlewares.Authenticate)
}
