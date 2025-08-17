package document

import (
	"github.com/google/uuid"
	"gofr.dev/pkg/gofr"
	"tnp-service/internal/models"
)

type DocServicer interface {
	CreateDoc(ctx *gofr.Context, data *models.DocumentRequest) error
	GetDoc(ctx *gofr.Context, docId *uuid.UUID) (*models.DocumentData, error)
}

type Handler struct {
	Service DocServicer
}

func New(ds DocServicer) *Handler {
	return &Handler{
		Service: ds,
	}
}

// Document handlers
func (h *Handler) CreateDocument(ctx *gofr.Context) (any, error) {
	return nil, nil
}

func (h *Handler) DeleteDocument(ctx *gofr.Context) (any, error) {
	return nil, nil
}

func (h *Handler) GetDocument(ctx *gofr.Context) (any, error) {
	return nil, nil
}
