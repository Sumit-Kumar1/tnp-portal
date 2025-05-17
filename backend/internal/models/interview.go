package models

import (
	"time"

	"github.com/google/uuid"
)

type InterviewData struct {
	ID            *uuid.UUID
	ApplicationID *uuid.UUID
	DateTime      time.Time
	Type          string // ONLINE or OFFLINE
	Location      string
	Notes         string
	Status        string // SCHEDULED, COMPLETED, CANCELLED
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type InterviewReq struct {
	ApplicationID string `json:"applicationId"`
	DateTime      string `json:"dateTime"`
	Type          string `json:"type"`
	Location      string `json:"location,omitempty"`
	Notes         string `json:"notes,omitempty"`
}

type InterviewResp struct {
	ID            *string   `json:"id,omitempty"`
	ApplicationID *string   `json:"applicationId,omitempty"`
	DateTime      *string   `json:"dateTime,omitempty"`
	Type          *string   `json:"type,omitempty"`
	Location      *string   `json:"location,omitempty"`
	Notes         *string   `json:"notes,omitempty"`
	Status        *string   `json:"status,omitempty"`
	CreatedAt     *DateTime `json:"createdAt,omitempty"`
	UpdatedAt     *DateTime `json:"updatedAt,omitempty"`
}
