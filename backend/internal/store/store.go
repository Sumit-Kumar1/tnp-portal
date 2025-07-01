package store

import "tnp-service/internal/models"

type Store struct {
}

func New() *Store {
	return &Store{}
}

func (s *Store) Get(id string) (*models.DocumentData, error) {
	return nil, nil
}

func (s *Store) GetAll() ([]*models.DocumentData, error) {
	return nil, nil
}

func (s *Store) Create(doc *models.DocumentData) error {
	return nil
}

func (s *Store) Update(doc *models.DocumentData) error {
	return nil
}

func (s *Store) Delete(id string) error {
	return nil
}
