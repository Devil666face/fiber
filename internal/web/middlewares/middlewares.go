package middlewares

import (
	"github.com/Devil666face/fiber/internal/web/handlers"

	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Logger(h *handlers.Handler) error {
	return logger.New()(h.Ctx())
}

func Recover(h *handlers.Handler) error {
	return recover.New()(h.Ctx())
}

func SecureHeaders(h *handlers.Handler) error {
	// XSSProtection:             "0",
	// ContentTypeNosniff:        "nosniff",
	// XFrameOptions:             "SAMEORIGIN",
	// ReferrerPolicy:            "no-referrer",
	// CrossOriginEmbedderPolicy: "require-corp",
	// CrossOriginOpenerPolicy:   "same-origin",
	// CrossOriginResourcePolicy: "same-origin",
	// OriginAgentCluster:        "?1",
	// XDNSPrefetchControl:       "off",
	// XDownloadOptions:          "noopen",
	// XPermittedCrossDomain:     "none",
	return helmet.New(helmet.Config{
		XSSProtection:  "1",
		ReferrerPolicy: "same-origin",
	})(h.Ctx())
}

func Compress(h *handlers.Handler) error {
	return compress.New()(h.Ctx())
}

func EncryptCookie(h *handlers.Handler) error {
	return encryptcookie.New(encryptcookie.Config{
		Key: h.Config().CookieKey,
	})(h.Ctx())
}

func Limiter(h *handlers.Handler) error {
	return limiter.New(limiter.Config{
		Storage: h.Storage(),
		Max:     h.Config().MaxQueryPerMinute,
	})(h.Ctx())
}
