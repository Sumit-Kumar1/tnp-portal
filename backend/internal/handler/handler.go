package handler

import (
	"gofr.dev/pkg/gofr"
)

type Handler struct {
	Service Servicer
}

func New(s Servicer) *Handler {
	return &Handler{Service: s}
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
