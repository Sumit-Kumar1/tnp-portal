package models

import (
	"time"

	"github.com/google/uuid"
)

type CompanyData struct {
	ID          *uuid.UUID
	Name        string
	Description string
	Website     string
	Industry    string
	Size        string
	Location    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type CompanyRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Website     string `json:"website,omitempty"`
	Industry    string `json:"industry"`
	Size        string `json:"size,omitempty"`
	Location    string `json:"location"`
}

type CompanyResponse struct {
	ID          *string   `json:"id,omitempty"`
	Name        *string   `json:"name,omitempty"`
	Description *string   `json:"description,omitempty"`
	Website     *string   `json:"website,omitempty"`
	Industry    *string   `json:"industry,omitempty"`
	Size        *string   `json:"size,omitempty"`
	Location    *string   `json:"location,omitempty"`
	CreatedAt   *DateTime `json:"createdAt,omitempty"`
	UpdatedAt   *DateTime `json:"updatedAt,omitempty"`
}
