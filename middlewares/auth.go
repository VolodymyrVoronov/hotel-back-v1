package middlewares

import (
	"hotel-back-v1/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")

	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	err := utils.ValidateToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	userId, err := utils.GetUserIDFromToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.Set("token", token)
	c.Set("userId", userId)

	c.Next()
}
