package middlewares

import (
	"github.com/Devil666face/fiber/internal/web/handlers"
	"github.com/Devil666face/fiber/internal/web/view"

	"github.com/gofiber/fiber/v2"
)

func Htmx(h *handlers.Handler) error {
	h.Ctx().Locals(view.Htmx, false)
	if _, ok := h.Ctx().GetReqHeaders()[view.HxRequest]; ok {
		h.Ctx().Locals(view.Htmx, true)
	}
	return h.Ctx().Next()
}

func HxOnly(h *handlers.Handler) error {
	if h.View().IsHtmx() {
		return h.Ctx().Next()
	}
	return fiber.ErrBadRequest
}
