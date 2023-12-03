package middlewares

import (
	"github.com/Devil666face/fiber/internal/models"
	"github.com/Devil666face/fiber/internal/web/handlers"

	"github.com/gofiber/fiber/v2"
)

var ErrNotPermissions = fiber.ErrNotFound

func Admin(h *handlers.Handler) error {
	if user, ok := h.ViewCtx().Locals(handlers.UserKey).(models.User); ok {
		if user.Admin {
			return h.ViewCtx().Next()
		}
	}
	return ErrNotPermissions
}
