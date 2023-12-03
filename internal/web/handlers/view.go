package handlers

import (
	"fmt"
	"log/slog"

	"github.com/gofiber/fiber/v2"
)

type ViewCtx struct {
	*fiber.Ctx
}

func NewViewCtx(c *fiber.Ctx) *ViewCtx {
	return &ViewCtx{c}
}

func (c ViewCtx) RenderWithCtx(name string, bind fiber.Map, layouts ...string) error {
	bind["c"] = c
	return c.Render(name, bind, layouts...)
}

func (c ViewCtx) CsrfToken() string {
	if token, ok := c.Locals(Csrf).(string); ok {
		return token
	}
	return ""
}

func (c ViewCtx) IsHtmx() bool {
	if htmx, ok := c.Locals(Htmx).(bool); ok {
		return htmx
	}
	return false
}

func (c ViewCtx) ClientRedirect(redirectURL string) error {
	c.Set(HXRedirect, redirectURL)
	return c.SendStatus(fiber.StatusFound)
}

func (c ViewCtx) URL(name string) string {
	return c.getRouteURL(name, fiber.Map{})
}

func (c ViewCtx) URLto(name, key string, val any) string {
	return c.getRouteURL(name, fiber.Map{
		key: val,
	})
}

func (c ViewCtx) getRouteURL(name string, fmap fiber.Map) string {
	url, err := c.GetRouteURL(name, fmap)
	if err != nil {
		slog.Error(fmt.Sprintf("Url %s not found", name))
	}
	return url
}

// func (c Ctx) Csrf() template.HTML {
// 	html := `<input type="hidden" name="csrf" value="%s">`
// 	//nolint:gosec //Because not revive data from user
// 	return template.HTML(fmt.Sprintf(html, c.CsrfToken()))
// }

// func (c ViewCtx) SetClientRefresh() {
// 	c.Set(HXRefresh, "true")
// }

// func (c ViewCtx) IsHtmxCurrentURL() bool {
// 	if url, ok := c.GetReqHeaders()[HxCurrentURL]; ok {
// 		return url[0] == c.BaseURL()+c.OriginalURL()
// 	}
// 	return false
// }
