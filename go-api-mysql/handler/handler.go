package handler

import (
	"go-api-docker/database"

	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{"data": "ok"})
}

func Get(c *gin.Context) {
	var schedules []database.Schedule
	database.DB.Find(&schedules)
	c.JSON(200, schedules)
}

func Create(c *gin.Context) {
	var schedule database.Schedule
	if err := c.ShouldBindJSON(&schedule); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&schedule)
	c.JSON(200, gin.H{"data": "Schedule is created successfully"})
}

func Update(c *gin.Context) {
	id := c.Param("id")
	var schedule database.Schedule
	if err := database.DB.First(&schedule, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Schedule not found"})
	}
	if err := c.ShouldBindJSON(&schedule); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	database.DB.Save(&schedule)
	c.JSON(200, gin.H{"data": "Schedule is updated successfully"})
}

func Delete(c *gin.Context) {
	id := c.Param("id")
	var schedule database.Schedule
	if database.DB.First(&schedule, id).Error != nil {
		c.JSON(404, gin.H{"error": "Schedule not found"})
		return
	}
	database.DB.Delete(&schedule)
	c.JSON(200, gin.H{"data": "Schedule is deleted successfully"})
}
