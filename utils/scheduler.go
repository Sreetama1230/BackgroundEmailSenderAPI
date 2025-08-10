package utils

import (
	"fmt"
	"new_start/BackgroundEmailSenderAPI/db"
	"new_start/BackgroundEmailSenderAPI/models"
	"time"
)

func StartEmailScheduler() {

	go func() {
		ticker := time.NewTicker(1 * time.Minute)
		defer ticker.Stop()
		for range ticker.C {
			var pendingEmails []models.Email

			db.DB.Where("status = ?", "pending").Find(&pendingEmails)
			for _, email := range pendingEmails {
				fmt.Printf("Sending email to %s: %s\n", email.Receiver, email.Subject)
				email.Status = "sent"
				email.ReceivedAt = time.Now()
				db.DB.Save(&email)
			}
			fmt.Println("Email scheduler ran at", time.Now())

		}
	}()
}
