package models

import (
	"github.com/google/uuid"
	"time"
)

type JobPosting struct {
	JobID          uuid.UUID
	EmployerName   string
	Description    string
	Location       string
	Qualifications string
	Salary         int
	docID          *uuid.UUID
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
