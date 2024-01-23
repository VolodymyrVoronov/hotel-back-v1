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

	isRoomAvailable, err := models.SearchAvailabilityByDatesByRoomID(booking.RoomID, booking.StartDate, booking.EndDate)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !isRoomAvailable {
		c.JSON(http.StatusOK, gin.H{"message": "Room not available", "booked": false})
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

	c.JSON(http.StatusOK, gin.H{"message": "Booking created successfully", "booked": true})
}

func GetAllBookings(c *gin.Context) {
	bookings, err := models.SelectAllBookings()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, bookings)
}

func ProcessBooking(c *gin.Context) {
	var processedBooking models.ProcessedBooking

	err := c.ShouldBindJSON(&processedBooking)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = processedBooking.UpdateBooking()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Booking updated successfully", "processed": true})
}

func GetSelectedBooking(c *gin.Context) {
	var booking models.BookedRoomID

	err := c.ShouldBindJSON(&booking)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	bookedRoom, err := models.SearchSelectedRoomByRoomID(booking.RoomID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, bookedRoom)
}

func CheckRoomAvailability(c *gin.Context) {
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
