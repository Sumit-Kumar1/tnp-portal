package documentstore

import (
	"gofr.dev/pkg/gofr"
	"tnp-service/internal/models"
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
