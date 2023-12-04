package middlewares

import (
	"github.com/Devil666face/fiber/internal/models"
	"github.com/Devil666face/fiber/internal/web/handlers"
	"github.com/Devil666face/fiber/internal/web/view"

	"github.com/gofiber/fiber/v2"
)

var ErrNotPermissions = fiber.ErrNotFound

func Admin(h *handlers.Handler) error {
	if user, ok := h.Ctx().Locals(view.UserKey).(models.User); ok {
		if user.Admin {
			return h.Ctx().Next()
		}
	}
	return ErrNotPermissions
}
