package utils

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/abhishekkushwahaa/AttendX/db"
	"github.com/abhishekkushwahaa/AttendX/models"
	"github.com/gin-gonic/gin"
	"github.com/go-webauthn/webauthn/webauthn"
	"go.mongodb.org/mongo-driver/bson"
)

var webAuthn *webauthn.WebAuthn
var sessionStore = make(map[string]*webauthn.SessionData)

func init() {
	config := &webauthn.Config{
		RPDisplayName: "AttendX",
		RPID:          "localhost",
		RPOrigins:     []string{"http://localhost:3000"},
	}
	var err error
	webAuthn, err = webauthn.New(config)
	if err != nil {
		log.Fatalf("Failed to initialize WebAuthn: %v", err)
	}
}

func BeginWebAuthnAuth(c *gin.Context) {
	username := c.Query("username")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username is required"})
		return
	}

	var user models.User
	err := db.UserColl.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	options, sessionData, err := webAuthn.BeginLogin(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to begin authentication"})
		return
	}

	sessionStore[username] = sessionData

	c.JSON(http.StatusOK, options)
}

func FinishWebAuthnAuth(c *gin.Context) {
	username := c.Query("username")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username is required"})
		return
	}

	var user models.User
	err := db.UserColl.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	sessionData, exists := sessionStore[username]
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid WebAuthn session"})
		return
	}

	_, err = webAuthn.FinishLogin(&user, *sessionData, c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}

	_, err = db.AttendColl.InsertOne(context.TODO(), bson.M{
		"username":  username,
		"timestamp": time.Now(),
		"method":    "fingerprint",
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to mark attendance"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Attendance marked with fingerprint"})
}
