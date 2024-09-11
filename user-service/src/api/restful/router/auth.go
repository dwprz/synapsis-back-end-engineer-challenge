package router

import (
	"user-service/src/api/restful/handler"
	"user-service/src/api/restful/middleware"

	"github.com/gofiber/fiber/v2"
)

func Auth(app *fiber.App, h *handler.Auth, m *middleware.Middleware) {
	// all
	app.Add("POST", "/api/auth/register", h.Register)
	app.Add("POST", "/api/auth/login", h.Login)
	app.Add("POST", "/api/auth/token/refresh", h.RefreshToken)
	app.Add("POST", "/api/auth/logout", h.Logout)
}