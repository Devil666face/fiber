package handlers

import (
	"github.com/Devil666face/fiber/internal/web/view"
	"github.com/gofiber/fiber/v2"
)

func DefaultErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	//nolint:errorlint //Because crash page for any errors, if not convertation to fiber.Error - return 500
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}
	c.Response().Header.SetContentType(fiber.MIMETextHTMLCharsetUTF8)
	return view.Error(
		code,
		err.Error(),
	).Render(c.UserContext(), c.Response().BodyWriter())
}
