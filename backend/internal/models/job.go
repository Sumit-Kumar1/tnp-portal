package models

import (
	"time"

	"github.com/google/uuid"
)

type JobData struct {
	ID           *uuid.UUID
	Title        string
	Description  string
	Requirements []string
	Location     string
	SalaryRange  struct {
		Min float64
		Max float64
	}
	Deadline  time.Time
	CompanyID *uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}

type JobRequest struct {
	Title        string   `json:"title"`
	Description  string   `json:"description"`
	Requirements []string `json:"requirements"`
	Location     string   `json:"location"`
	SalaryRange  struct {
		Min float64 `json:"min"`
		Max float64 `json:"max"`
	} `json:"salaryRange"`
	Deadline string `json:"deadline"`
}

type JobResponse struct {
	ID           *string  `json:"id,omitempty"`
	Title        *string  `json:"title,omitempty"`
	Description  *string  `json:"description,omitempty"`
	Requirements []string `json:"requirements,omitempty"`
	Location     *string  `json:"location,omitempty"`
	SalaryRange  struct {
		Min *float64 `json:"min,omitempty"`
		Max *float64 `json:"max,omitempty"`
	} `json:"salaryRange,omitempty"`
	Deadline  *string   `json:"deadline,omitempty"`
	CompanyID *string   `json:"companyId,omitempty"`
	CreatedAt *DateTime `json:"createdAt,omitempty"`
	UpdatedAt *DateTime `json:"updatedAt,omitempty"`
}
