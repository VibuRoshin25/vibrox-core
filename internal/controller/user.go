package controller

import (
	"net/http"

	"vibrox-core/internal/config"
	"vibrox-core/internal/logs"
	"vibrox-core/internal/models"

	"github.com/gin-gonic/gin"
)

// GetHealth validates server health
func GetHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "User Service is up and running",
	})
}

// GetUsers lists users
func GetUsers(c *gin.Context) {
	users := []models.User{}
	config.DB.Find(&users)
	if len(users) == 0 {
		logs.LogError(c, "No users available in the database")
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No Users Available",
		})
		return
	}
	c.JSON(http.StatusOK, &users)
}

// GetById gets specified user by id
func GetById(c *gin.Context) {
	var user models.User
	config.DB.Where("id = ?", c.Param("id")).Find(&user)
	if user.CreatedAt.IsZero() {
		logs.LogError(c, "User not found with ID: "+c.Param("id"))
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User Not Found",
		})
		return
	}
	c.JSON(http.StatusOK, &user)
}

// CreateUser creates user
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		logs.LogError(c, "Invalid create user request: "+err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid create user request"})
		return
	}
	if err := config.DB.Create(&user).Error; err != nil {
		logs.LogError(c, "Failed to create user: "+err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	c.JSON(http.StatusOK, &user)
}

// DeleteUser deletes user
func DeleteUser(c *gin.Context) {
	var user models.User
	if err := config.DB.Where("id = ?", c.Param("id")).Delete(&user).Error; err != nil {
		logs.LogError(c, "Failed to delete user: "+err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}
	c.JSON(http.StatusOK, &user)
}

// UpdateUser edits user profile
func UpdateUser(c *gin.Context) {
	var user models.User
	if err := config.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		logs.LogError(c, "Failed to find user: "+err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find user"})
		return
	}
	if err := c.BindJSON(&user); err != nil {
		logs.LogError(c, "Invalid update user request: "+err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid update user request"})
		return
	}
	if err := config.DB.Save(&user).Error; err != nil {
		logs.LogError(c, "Failed to update user: "+err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}
	c.JSON(http.StatusOK, &user)
}
