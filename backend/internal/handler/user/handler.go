package user

import "gofr.dev/pkg/gofr"

type Handler struct {
}

func New() *Handler {
	return &Handler{}
}

// User Endpoints
func (h *Handler) CreateUser(ctx *gofr.Context) (any, error) {
	return nil, nil
}
func (h *Handler) UpdateUser(ctx *gofr.Context) (any, error) {
	return nil, nil
}

func (h *Handler) DeleteUser(ctx *gofr.Context) (any, error) {
	return nil, nil
}
