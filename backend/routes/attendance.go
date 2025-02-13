package routes

import (
	"github.com/abhishekkushwahaa/AttendX/utils"
	"github.com/gin-gonic/gin"
)

func AttendanceRoutes(r *gin.Engine) {
	r.POST("/attendance/fingerprint", utils.BeginWebAuthnAuth)
	r.POST("/attendance/fingerprint/verify", utils.FinishWebAuthnAuth)
}
