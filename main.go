package main

import (
	"new_start/BackgroundEmailSenderAPI/auth"
	"new_start/BackgroundEmailSenderAPI/db"
	"new_start/BackgroundEmailSenderAPI/handlers"
	"new_start/BackgroundEmailSenderAPI/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	utils.StartEmailScheduler()
	r := gin.Default()
	r.POST("/login", handlers.Login)       // public
	r.POST("/register", handlers.Register) // public
	authorized := r.Group("/")
	authorized.Use(auth.AuthMiddleware())
	{
		authorized.POST("/schedule-email", handlers.ScheduleEmail)
		// other protected routes
	}

	r.Run(":8080") // Run on port 8080
}
