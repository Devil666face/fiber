package handlers

import (
	"github.com/Devil666face/fiber/internal/web/view"
	"github.com/gofiber/fiber/v2"
)

func Index(h *Handler) error {
	return h.RenderTempl(view.Index, fiber.Map{})
}
