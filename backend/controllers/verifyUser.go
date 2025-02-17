package controllers

import (
	"context"
	"net/http"

	"github.com/abhishekkushwahaa/AttendX/db"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func VerifyUser(c *gin.Context) {
	var req struct {
		Username    string `json:"username"`
		Fingerprint string `json:"fingerprint"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	var user bson.M
	err := db.DB.Collection("users").FindOne(context.TODO(), bson.M{"username": req.Username}).Decode(&user)
	if err != nil || user["fingerprint"] != req.Fingerprint {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Fingerprint verification failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User verified successfully"})
}
