package router

import (
	"user-service/src/api/restful/handler"
	"user-service/src/api/restful/middleware"

	"github.com/gofiber/fiber/v2"
)

func User(app *fiber.App, uh *handler.User, m *middleware.Middleware) {
	app.Add("GET", "/api/users/current", m.VerifyJwt, uh.GetByCurrent)
	app.Add("PATCH", "/api/users/current", m.VerifyJwt, uh.Update)
	app.Add("DELETE", "/api/users/:userId", m.VerifyJwt, m.VerifyAdmin, uh.Delete)
}
