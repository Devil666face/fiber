package routes

import (
	"github.com/Devil666face/fiber/internal/web/handlers"
	"github.com/Devil666face/fiber/internal/web/middlewares"
)

func (r *Router) setIndex() {
	index := r.router.Group("/")
	index.Use(r.wrapper(middlewares.Auth))

	index.Get(
		"",
		r.wrapper(handlers.Index),
	).Name("index")
}
