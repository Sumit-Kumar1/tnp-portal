package models

import (
	"github.com/google/uuid"
	"time"
)

type StudentData struct {
	StudentID   *uuid.UUID
	UserID      *uuid.UUID
	Name        *string
	Degree      *string
	Department  *string
	Major       *string
	GradYear    *int
	ContactInfo *string
	Education   []string
	Skills      []string
	docID       *uuid.UUID
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}

type StudentReq struct {
	Name       string `json:"name"`
	Degree     string `json:"degree"`
	Department string `json:"department"`
	Major      string `json:"major"`
	GradYear   string `json:"graduationYear"`
	Contact    string `json:"contactInfo"`
	Education  string `json:"educationBackground"`
	Skill      string `json:"skills"`
	DocId      string `json:"docId"`
}

type StudentResp struct {
	ID         *string   `json:"id,omitempty"`
	UserID     *string   `json:"userId,omitempty"`
	Name       *string   `json:"name,omitempty"`
	Degree     *string   `json:"degree,omitempty"`
	Department *string   `json:"department,omitempty"`
	Major      *string   `json:"major,omitempty"`
	GradYear   *string   `json:"graduationYear,omitempty"`
	Contact    *string   `json:"contactInfo,omitempty"`
	Education  *string   `json:"educationBackground,omitempty"`
	Skill      *string   `json:"skills,omitempty"`
	DocId      *string   `json:"docId,omitempty"`
	CreatedAt  *DateTime `json:"createdAt,omitempty"`
	UpdatedAt  *DateTime `json:"updatedAt,omitempty"`
}
