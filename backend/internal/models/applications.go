package models

import (
	"time"

	"github.com/google/uuid"
)

type ApplicationData struct {
	ApplicationID *uuid.UUID
	JobID         *uuid.UUID
	StudentID     *uuid.UUID
	Status        string
	AppliedAt     time.Time
	UpdatedAt     time.Time
}
