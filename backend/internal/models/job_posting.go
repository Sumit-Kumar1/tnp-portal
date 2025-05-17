package models

import (
	"time"

	"github.com/google/uuid"
)

type JobPosting struct {
	JobID          uuid.UUID
	EmployerName   string
	Description    string
	Location       string
	Qualifications string
	Salary         int
	DocID          *uuid.UUID
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
