package documentsvc

import (
	"tnp-service/internal/models"

	"github.com/google/uuid"
	"gofr.dev/pkg/gofr"
)

type DocStorer interface {
	Create(ctx *gofr.Context, data *models.DocumentData) error
	Get(ctx *gofr.Context) (*models.DocumentData, error)
}

type DocService struct {
	Store DocStorer
}

func New(ds DocStorer) *DocService {
	return &DocService{
		Store: ds,
	}
}

func (ds *DocService) CreateDoc(ctx *gofr.Context, data *models.DocumentReq) error {
	return nil
}

func (ds *DocService) GetDoc(ctx *gofr.Context, docId *uuid.UUID) (*models.DocumentData, error) {
	return nil, nil
}
