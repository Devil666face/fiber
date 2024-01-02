package view

import (
	"fmt"
	"log/slog"

	"github.com/gofiber/fiber/v2"
)

const (
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
	AuthKey    = "authenticated"
	UserKey    = "User"
	UsersKey   = "Users"
	MessageKey = "Message"
	SuccessKey = "Success"
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

func (c View) IsURL(name string) bool {
	return c.Ctx.OriginalURL() == c.URL(name)
}

func (c View) getRouteURL(name string, fmap fiber.Map) string {
	url, err := c.GetRouteURL(name, fmap)
	if err != nil {
		slog.Error(fmt.Sprintf("url %s not found", name))
	}
	return url
}

func (c View) IsHtmx() bool {
	if htmx, ok := c.Locals(Htmx).(bool); ok {
		return htmx
	}
	return false
}

func (c View) ClientRedirect(redirectURL string) error {
	c.Set(HXRedirect, redirectURL)
	return c.SendStatus(fiber.StatusFound)
}

// func (c ViewCtx) SetClientRefresh() {
// 	c.Set(HXRefresh, "true")
// }

// func (c ViewCtx) IsHtmxCurrentURL() bool {
// 	if url, ok := c.GetReqHeaders()[HxCurrentURL]; ok {
// 		return url[0] == c.BaseURL()+c.OriginalURL()
// 	}
// 	return false
// }
