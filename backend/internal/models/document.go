package models

import (
	"time"

	"github.com/google/uuid"
)

type DocumentData struct {
	DocID     *uuid.UUID
	UserID    *uuid.UUID
	DocType   *string
	Document  []byte
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

type DocumentRequest struct {
	UserID   string `json:"userId"`
	DocName  string `json:"docName"`
	DocType  string `json:"docType"`
	Document []byte `json:"document"`
}

type DocumentResponse struct {
	DocID     *string   `json:"docId,omitempty"`
	UserID    *string   `json:"userId,omitempty"`
	DocName   *string   `json:"docName,omitempty"`
	DocType   *string   `json:"docType,omitempty"`
	Document  []byte    `json:"document,omitempty"`
	CreatedAt *DateTime `json:"createdAt,omitempty"`
	UpdatedAt *DateTime `json:"updatedAt,omitempty"`
}
