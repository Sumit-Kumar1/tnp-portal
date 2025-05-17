package models

import (
	"time"

	"github.com/google/uuid"
)

type NotificationData struct {
	ID        *uuid.UUID
	UserID    *uuid.UUID
	Title     string
	Message   string
	Type      string // JOB, INTERVIEW, APPLICATION, SYSTEM
	IsRead    bool
	CreatedAt time.Time
}

type NotificationResponse struct {
	ID        *string   `json:"id,omitempty"`
	UserID    *string   `json:"userId,omitempty"`
	Title     *string   `json:"title,omitempty"`
	Message   *string   `json:"message,omitempty"`
	Type      *string   `json:"type,omitempty"`
	IsRead    *bool     `json:"isRead,omitempty"`
	CreatedAt *DateTime `json:"createdAt,omitempty"`
}
