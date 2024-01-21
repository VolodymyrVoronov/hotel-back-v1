package routes

import (
	"hotel-back-v1/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.POST("/bookings", CreateBooking)
	server.GET("/bookings", GetAllBookings)
	server.POST("/bookings-room", GetSelectedBooking)

	server.POST("/availability", CheckRoomAvailability)

	server.POST("/contact-us", ContactUs)
	server.POST("/subscription", RegisterSubscription)

	server.POST("/login", Login)

	authenticated := server.Group("/admin")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/register", Register)
	authenticated.POST("/roles", CreateRole)
	authenticated.GET("/roles", GetRoles)
}
