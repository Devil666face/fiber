package view

import (
	"fmt"
	"log/slog"

	"github.com/gofiber/fiber/v2"
)

const (
	AuthKey      = "authenticated"
	UserKey      = "User"
	UserID       = "user_id"
	Csrf         = "csrf"
	Htmx         = "htmx"
	HxRequest    = "Hx-Request"
	HxCurrentURL = "Hx-Current-Url"
	HXRedirect   = "HX-Redirect"
	Host         = "Host"
	// HXRefresh    = "HX-Refresh"
)

type View struct {
	*fiber.Ctx
}

func New(c *fiber.Ctx) *View {
	return &View{c}
}

func (c View) CsrfToken() string {
	if token, ok := c.Locals(Csrf).(string); ok {
		return token
	}
	return ""
}

func (c View) URL(name string) string {
	return c.getRouteURL(name, fiber.Map{})
}

func (c View) URLto(name, key string, val any) string {
	return c.getRouteURL(name, fiber.Map{
		key: val,
	})
}

func (c View) getRouteURL(name string, fmap fiber.Map) string {
	url, err := c.GetRouteURL(name, fmap)
	if err != nil {
		slog.Error(fmt.Sprintf("Url %s not found", name))
	}
	return url
}
