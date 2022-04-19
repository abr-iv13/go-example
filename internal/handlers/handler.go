package handlers

import (
	"project/back/internal/services"
	"project/back/pkg/logger"
)

type Handler struct {
	Services *services.Services
	log      *logger.Logger
}

func NewHandler(services *services.Services, log *logger.Logger) *Handler {
	return &Handler{Services: services}
}
