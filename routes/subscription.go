package routes

import (
	"fmt"
	"hotel-back-v1/models"
	"hotel-back-v1/services"
	"net/http"
	"net/mail"

	"github.com/gin-gonic/gin"
)

func GetAllSubscriptions(c *gin.Context) {
	subscriptions, err := models.SelectAllSubscriptions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, subscriptions)
}

func RegisterSubscription(c *gin.Context) {
	var subscription models.Subscription

	err := c.ShouldBindJSON(&subscription)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	isSubscriptionExisted, err := models.CheckSubscription(subscription.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if isSubscriptionExisted {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "You are already subscribed to our newsletter"})
		return
	}

	err = services.SendMail("127.0.0.1:1025", (&mail.Address{Name: fmt.Sprintf("User %s", subscription.Email), Address: subscription.Email}).String(), "Subscription", fmt.Sprintf("User with email %s subscribed to our newsletter", subscription.Email), []string{(&mail.Address{Name: "Admin", Address: "to@example.com"}).String()})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = subscription.InsertSubscription()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Subscription created successfully"})
}
