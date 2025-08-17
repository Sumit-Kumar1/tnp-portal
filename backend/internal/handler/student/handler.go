package student

import "gofr.dev/pkg/gofr"

type Handler struct {
}

func New() *Handler {
	return &Handler{}
}

// student methods
func (h *Handler) CreateStudent(ctx *gofr.Context) (any, error) {
	return nil, nil
}

func (h *Handler) GetStudent(ctx *gofr.Context) (any, error) {
	return nil, nil
}

func (h *Handler) UpdateStudent(ctx *gofr.Context) (any, error) {
	return nil, nil
}

func (h *Handler) DeleteStudent(ctx *gofr.Context) (any, error) {
	return nil, nil
}
