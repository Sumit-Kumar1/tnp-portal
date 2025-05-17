package documenthttp

import (
	"tnp-backend/internal/models"

	"github.com/google/uuid"
	"gofr.dev/pkg/gofr"
)

type DocServicer interface {
	CreateDoc(ctx *gofr.Context, data *models.DocumentRequest) error
	GetDoc(ctx *gofr.Context, docId *uuid.UUID) (*models.DocumentData, error)
}

type DocHandler struct {
	Service DocServicer
}

func New(ds DocServicer) *DocHandler {
	return &DocHandler{
		Service: ds,
	}
}

func (dh *DocHandler) CreateDoc(ctx *gofr.Context) (any, error) {
	return nil, nil
}

func (dh *DocHandler) GetDoc(ctx *gofr.Context) (any, error) {
	return nil, nil
}
