package routes

import (
	"hotel-back-v1/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.POST("/bookings", CreateBooking)
	server.GET("/bookings", GetBookings)
	server.POST("/bookings-room", GetBooking)

	server.POST("/availability", CheckAvailability)

	server.POST("/roles", CreateRole)
	server.GET("/roles", GetRoles)

	server.POST("/register", Register)
	server.POST("/login", Login)

	authenticated := server.Group("/admin")
	authenticated.Use(middlewares.Authenticate)
}
