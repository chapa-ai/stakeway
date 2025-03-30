package handler

import (
	"github.com/gofiber/fiber/v2"
	"stakeway/internal/service"
)

func New(s service.Service) *fiber.App {
	app := fiber.New(fiber.Config{})

	h := &Handler{
		service: s,
	}

	api := app.Group("/api")
	{
		api.Post("/validators", h.CreateValidators)
		api.Get("/validators/:request_id", h.GetValidatorStatus)
	}

	return app
}
