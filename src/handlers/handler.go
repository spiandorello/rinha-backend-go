package handlers

import (
	"github.com/gofiber/fiber/v2"
	"rinha-de-backend/src/interfaces"
)

type Handlers struct {
	handlers []interfaces.Handler
}

func New(handlers []interfaces.Handler) *Handlers {
	return &Handlers{
		handlers: handlers,
	}
}

func (h *Handlers) GetRoutes(app *fiber.App) {
	for _, handler := range h.handlers {
		handler.Routes(app)
	}
}

func Routes(app *fiber.App, h *Handlers) {
	h.GetRoutes(app)
}
