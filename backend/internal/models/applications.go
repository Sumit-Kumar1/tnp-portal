package models

import (
	"github.com/google/uuid"
	"time"
)

type ApplicationData struct {
	ApplicationID *uuid.UUID
	JobID         *uuid.UUID
	StudentID     *uuid.UUID
	Status        string
	AppliedAt     time.Time
	UpdatedAt     time.Time
}
