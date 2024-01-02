package middlewares

import (
	"github.com/Devil666face/fiber/internal/web/handlers"
	"github.com/Devil666face/fiber/internal/web/view"

	"github.com/gofiber/fiber/v2/middleware/csrf"
)

func Csrf(h *handlers.Handler) error {
	// KeyLookup:         "header:" + "X-Csrf-Token",
	// CookieName:        "csrf_",
	// CookieSameSite:    "Lax",
	// Expiration:        1 * time.Hour,
	// KeyGenerator:      utils.UUIDv4,
	// ErrorHandler:      defaultErrorHandler,
	// Extractor:         CsrfFromHeader("X-Csrf-Token"),
	// SessionKey:        "fiber.csrf.token",
	// HandlerContextKey: "fiber.csrf.handler",
	return csrf.New(csrf.Config{
		Storage: h.Storage(),
		// KeyLookup:  "form:csrf",
		ContextKey:     view.Csrf,
		CookieHTTPOnly: true,
		SingleUseToken: true,
	})(h.Ctx())
}
