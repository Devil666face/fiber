package view

import (
	"fmt"
	"log/slog"

	"github.com/Devil666face/fiber/internal/models"
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

const (
	userKey = "User"
)

type View struct {
	*fiber.Ctx
}

type Map fiber.Map

func (m Map) get(key string) any {
	if val, ok := m[key]; ok {
		return val
	}
	return fmt.Errorf("not found value for key: %s in view map", key)
}

func (m Map) getUser() models.User {
	if user, ok := m.get(userKey).(models.User); ok {
		return user
	}
	return models.User{}
}

func (m Map) notUser() bool {
	return (m.getUser() == models.User{})
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
