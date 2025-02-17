package controllers

import (
	"context"
	"net/http"

	"github.com/abhishekkushwahaa/AttendX/db"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rekognition"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func FaceRecognition(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Image    []byte `json:"image"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	var user bson.M
	err := db.DB.Collection("users").FindOne(context.TODO(), bson.M{"username": req.Username}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	sess := session.Must(session.NewSession(&aws.Config{Region: aws.String("us-east-1")}))
	rekognitionClient := rekognition.New(sess)

	input := &rekognition.CompareFacesInput{
		SourceImage: &rekognition.Image{Bytes: req.Image},
		TargetImage: &rekognition.Image{Bytes: user["face_image"].([]byte)},
	}

	result, err := rekognitionClient.CompareFaces(input)
	if err != nil || len(result.FaceMatches) == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Face recognition failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Face matched successfully"})
}
