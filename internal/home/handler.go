package home

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type HomeHandler struct {
	router       fiber.Router
	CustomLogger *zerolog.Logger
}

func NewHomeHandler(r fiber.Router, l *zerolog.Logger) {
	h := HomeHandler{
		router:       r,
		CustomLogger: l,
	}
	app := h.router.Group("/api")
	app.Get("/", h.Home)
	app.Get("/error", h.Error)
}

func (h *HomeHandler) Home(c *fiber.Ctx) error {

	return fiber.ErrBadRequest
}
func (h *HomeHandler) Error(c *fiber.Ctx) error {

	h.CustomLogger.Info().Msg("csadvfgb")
	return c.SendString("Error")
}
