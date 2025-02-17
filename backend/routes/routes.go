package routes

import (
	"github.com/abhishekkushwahaa/AttendX/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/login", controllers.LoginUser)
	r.POST("/register", controllers.RegisterUser)
	r.POST("/verify", controllers.VerifyUser)
	r.POST("/face-recognition", controllers.FaceRecognition)
	r.POST("/mark-attendance", controllers.MarkAttendance)
}
