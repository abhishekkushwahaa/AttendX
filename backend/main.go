package main

import (
	"github.com/abhishekkushwahaa/AttendX/config"
	"github.com/abhishekkushwahaa/AttendX/db"
	"github.com/abhishekkushwahaa/AttendX/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	db.InitDB()

	r := gin.Default()
	routes.RegisterRoutes(r)
	r.Run(":8081")
}
