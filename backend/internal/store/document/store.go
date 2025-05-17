package documentstore

import (
	"tnp-backend/internal/models"

	"gofr.dev/pkg/gofr"
)

type DocStore struct {
}

func New() *DocStore {
	return &DocStore{}
}

func (d *DocStore) Create(ctx *gofr.Context, data *models.DocumentData) error {
	return nil
}

func (d *DocStore) Get(ctx *gofr.Context) (*models.DocumentData, error) {
	return nil, nil
}
