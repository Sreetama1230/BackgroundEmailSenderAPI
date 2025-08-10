package handlers

import (
	"net/http"
	"new_start/BackgroundEmailSenderAPI/db"
	"new_start/BackgroundEmailSenderAPI/models"
	"time"

	"gorm.io/datatypes"

	"github.com/gin-gonic/gin"
)

type ScheduleEmailRequest struct {
	Subject      string                      `json:"subject"`
	Body         string                      `json:"body"`
	Sender       string                      `json:"sender"`
	Receiver     string                      `json:"receiver"`
	Attachments  datatypes.JSONSlice[string] // List of attachment file paths
	ScheduleTime time.Time                   `json:"schedule_time"` // ISO 8601 format
	IsDraft      bool                        `json:"is_draft"`      // true if the email is a draft
}

func ScheduleEmail(c *gin.Context) {

	var input ScheduleEmailRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": "Invalid input"})
		return
	}

	email := models.Email{
		Subject:     input.Subject,
		Body:        input.Body,
		Sender:      input.Sender,
		Receiver:    input.Receiver,
		Attachments: input.Attachments,
		ReceivedAt:  time.Now(),
		SentAt:      input.ScheduleTime,
		IsDraft:     input.IsDraft,
		Status:      "pending",
	}
	//save to db
	if result := db.DB.Create(&email); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to schedule email"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Email scheduled successfully",
	})
}
