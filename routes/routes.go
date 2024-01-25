package routes

import (
	"hotel-back-v1/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.POST("/bookings", CreateBooking)
	server.POST("/bookings-room", GetSelectedBooking)

	server.POST("/availability", CheckRoomAvailability)

	server.POST("/contact-us", ContactUs)
	server.POST("/subscription", RegisterSubscription)

	server.POST("/login", Login)

	authenticated := server.Group("/admin")
	authenticated.Use(middlewares.Authenticate)

	authenticated.POST("/register", Register)

	authenticated.GET("/bookings", GetAllBookings)
	authenticated.POST("/bookings-process", ProcessBooking)

	authenticated.GET("/contact-us", GetAllContactUs)
	authenticated.GET("/subscription", GetAllSubscriptions)

	authenticated.POST("/roles", CreateRole)
	authenticated.GET("/roles", GetRoles)

	authenticated.GET("/users", GetAllUsers)
	authenticated.DELETE("/users", DeleteUser)
}
