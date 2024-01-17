package routes

import (
	"hotel-back-v1/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBooking(c *gin.Context) {
	var booking models.Booking

	err := c.ShouldBindJSON(&booking)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = booking.InsertBooking()
	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var bookedRoom models.BookedRoom

	bookedRoom.RoomID = booking.RoomID
	bookedRoom.StartDate = booking.StartDate
	bookedRoom.EndDate = booking.EndDate

	err = bookedRoom.InsertBookedRoom()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Booking created successfully"})
}

func GetBookings(c *gin.Context) {
	bookings, err := models.SelectAllBookings()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, bookings)
}

func CheckAvailability(c *gin.Context) {
	var roomAvailability models.RoomAvailability

	err := c.ShouldBindJSON(&roomAvailability)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	available, err := models.SearchAvailabilityByDatesByRoomID(roomAvailability.RoomID, roomAvailability.StartDate, roomAvailability.EndDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "available": available})
		return
	}

	c.JSON(http.StatusOK, gin.H{"available": available})
}
