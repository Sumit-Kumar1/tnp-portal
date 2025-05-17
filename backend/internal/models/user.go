package models

import (
	"time"

	"github.com/google/uuid"
)

// UserData represents the user data in the database
type UserData struct {
	UserID       uuid.UUID `json:"user_id" db:"id"`
	EnrollmentID string    `json:"enrollment_id" db:"enrollment_id"`
	Name         string    `json:"name" db:"name"`
	Email        string    `json:"email" db:"email"`
	PasswordHash string    `json:"-" db:"password_hash"`
	Role         string    `json:"role" db:"role"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

// UserRequest represents the user request data
type UserRequest struct {
	EnrollmentID string `json:"enrollment_id" binding:"required"`
	Name         string `json:"name" binding:"required"`
	Email        string `json:"email" binding:"required,email"`
	Password     string `json:"password" binding:"required,min=8"`
	Role         string `json:"role" binding:"required,oneof=student admin company"`
}

// UserResponse represents the user response data
type UserResponse struct {
	UserID       uuid.UUID `json:"user_id"`
	EnrollmentID string    `json:"enrollment_id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Role         string    `json:"role"`
	Token        string    `json:"token,omitempty"`
}

// LoginRequest represents the login request data
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse represents the response format for login
type LoginResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}

// ConvertToUserResp converts UserData to UserResponse
func ConvertToUserResp(user *UserData) *UserResponse {
	return &UserResponse{
		UserID:       user.UserID,
		EnrollmentID: user.EnrollmentID,
		Name:         user.Name,
		Email:        user.Email,
		Role:         user.Role,
		Token:        user.PasswordHash,
	}
}
