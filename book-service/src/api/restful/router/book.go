package router

import (
	"book-service/src/api/restful/handler"
	"book-service/src/api/restful/middleware"

	"github.com/gofiber/fiber/v2"
)

func Book(app *fiber.App, bh *handler.Book, m *middleware.Middleware) {
	app.Add("POST", "/api/books", m.VerifyJwt, m.VerifyAdmin, bh.Add)
	app.Add("GET", "/api/books", m.VerifyJwt, bh.Get)
	app.Add("GET", "/api/books/recommendations", m.VerifyJwt, bh.GetRecommendations)
	app.Add("PATCH", "/api/books/:bookId", m.VerifyJwt, m.VerifyAdmin, bh.Update)
	app.Add("DELETE", "/api/books/:bookId", m.VerifyJwt, m.VerifyAdmin, bh.Delete)
}
