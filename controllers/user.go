package controllers

import (
	"gin-template/internal/models"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []models.User

	if err := models.DB.Find(&users).Error; err != nil {
		c.JSON(400, gin.H{"error": err})
	}

	c.JSON(200, users)
}
