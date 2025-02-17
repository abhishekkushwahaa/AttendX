package controllers

import (
	"context"
	"time"

	"github.com/abhishekkushwahaa/AttendX/db"
	"github.com/abhishekkushwahaa/AttendX/models"
	"github.com/gin-gonic/gin"
)

func MarkAttendance(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	attendance := models.Attendance{
		UserID:    req.Username,
		Timestamp: time.Now(),
		Status:    true,
	}

	_, err := db.DB.Collection("attendance").InsertOne(context.TODO(), attendance)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to mark attendance"})
		return
	}

	c.JSON(200, gin.H{"message": "Attendance marked successfully"})
}
