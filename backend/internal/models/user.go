package models

import (
	"github.com/google/uuid"
	"time"
)

type UserData struct {
	UserID       *uuid.UUID
	EnrollmentID string
	Name         string
	Email        string
	PasswordHash []byte
	Role         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type UserReq struct {
	Name             *string `json:"name,omitempty"`
	Email            *string `json:"email,omitempty"`
	Password         *string `json:"password,omitempty"`
	EnrollmentNumber *string `json:"enrollmentNumber,omitempty"`
}

type UserResp struct {
	UserID    *string `json:"userId,omitempty"`
	Name      *string `json:"name,omitempty"`
	Email     *string `json:"email,omitempty"`
	Role      *string `json:"role,omitempty"`
	CreatedAt *string `json:"createdAt,omitempty"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
}
