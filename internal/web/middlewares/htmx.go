package middlewares

import (
	"github.com/Devil666face/fiber/internal/web/handlers"
	"github.com/Devil666face/fiber/internal/web/view"

	"github.com/gofiber/fiber/v2"
)

func Htmx(h *handlers.Handler) error {
	h.ViewCtx().Locals(view.Htmx, false)
	if _, ok := h.ViewCtx().GetReqHeaders()[view.HxRequest]; ok {
		h.ViewCtx().Locals(view.Htmx, true)
	}
	return h.ViewCtx().Next()
}

func HxOnly(h *handlers.Handler) error {
	if h.ViewCtx().IsHtmx() {
		return h.ViewCtx().Next()
	}
	return fiber.ErrBadRequest
}
