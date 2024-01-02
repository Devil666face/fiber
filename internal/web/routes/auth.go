package routes

import (
	"github.com/Devil666face/fiber/internal/web/handlers"
	"github.com/Devil666face/fiber/internal/web/middlewares"
)

func (r *Router) setAuth() {
	auth := r.router.Group("/auth")

	auth.Get(
		"/login",
		r.wrapper(middlewares.AlreadyLogin),
		r.wrapper(handlers.LoginPage),
	).Name("login")
	auth.Post(
		"/login",
		r.wrapper(handlers.Login),
	)
}
