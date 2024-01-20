package routes

import (
	"hotel-back-v1/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateRole(c *gin.Context) {
	var role models.Role

	err := c.ShouldBindJSON(&role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = role.InsertRole()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Role created successfully"})
}

func GetRoles(c *gin.Context) {
	roles, err := models.SelectAllRoles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, roles)
}
