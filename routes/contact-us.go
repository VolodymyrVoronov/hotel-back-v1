package routes

import (
	"hotel-back-v1/models"
	"hotel-back-v1/services"
	"net/http"
	"net/mail"

	"github.com/gin-gonic/gin"
)

func ContactUs(c *gin.Context) {
	var contactUs models.ContactUs

	err := c.ShouldBindJSON(&contactUs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = services.SendMail("127.0.0.1:1025", (&mail.Address{Name: contactUs.Name, Address: contactUs.Email}).String(), "Contact Us", contactUs.Message, []string{(&mail.Address{Name: "Admin", Address: "to@example.com"}).String()})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = contactUs.InsertContactUs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Email sent successfully"})
}
