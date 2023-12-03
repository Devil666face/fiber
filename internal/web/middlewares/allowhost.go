package middlewares

import (
	"strings"

	"github.com/Devil666face/fiber/internal/web/handlers"
	"github.com/Devil666face/fiber/internal/web/view"

	"github.com/gofiber/fiber/v2"
)

func AllowHost(h *handlers.Handler) error {
	if host, ok := h.ViewCtx().GetReqHeaders()[view.Host]; ok {
		if strings.Contains(host[0], h.Config().AllowHost) {
			return h.ViewCtx().Next()
		}
	}
	return fiber.ErrBadRequest
}
