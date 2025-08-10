package models

import (
	"time"

	"gorm.io/datatypes"
)

type Email struct {
	ID          uint `gorm:"primaryKey"`
	Subject     string
	Body        string
	Sender      string
	Receiver    string
	SentAt      time.Time
	ReceivedAt  time.Time
	Attachments datatypes.JSONSlice[string] // List of attachment file paths
	IsDraft     bool
	Status      string // e.g., "sent", "received", "draft"

}
