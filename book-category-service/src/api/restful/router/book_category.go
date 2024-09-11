package router

import (
	"book-category-service/src/api/restful/handler"
	"book-category-service/src/api/restful/middleware"

	"github.com/gofiber/fiber/v2"
)

func BookCategory(app *fiber.App, bh *handler.BookCategory, m *middleware.Middleware) {
	// admin
	app.Add("POST", "/api/book-categories", m.VerifyJwt, m.VerifyAdmin, bh.Create)
	app.Add("DELETE", "/api/book-categories/:category/categories", m.VerifyJwt, m.VerifyAdmin, bh.DeleteByCategory)
	app.Add("DELETE", "/api/book-categories/:bookId/books", m.VerifyJwt, m.VerifyAdmin, bh.DeleteByBookId)
	app.Add("DELETE", "/api/book-categories", m.VerifyJwt, m.VerifyAdmin, bh.DeleteFromBookCategory)

	// all
	app.Add("GET", "/api/book-categories/:category", m.VerifyJwt, bh.GetByCategory)
}
