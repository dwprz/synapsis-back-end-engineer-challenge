package server

import (
	"book-category-service/src/api/restful/handler"
	"book-category-service/src/api/restful/middleware"
	"book-category-service/src/api/restful/router"
	"book-category-service/src/infrastructure/config"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// this main restful server
type Restful struct {
	app *fiber.App
}

func NewRestful(bh *handler.BookCategory) *Restful {
	m := middleware.New()

	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		IdleTimeout:   20 * time.Second,
		ReadTimeout:   20 * time.Second,
		WriteTimeout:  20 * time.Second,
		ErrorHandler:  m.Error,
	})

	app.Use(recover.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://restful.local:80",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:     "*",
		AllowCredentials: true,
	}))

	router.BookCategory(app, bh, m)

	return &Restful{
		app: app,
	}
}

func (r *Restful) Run() {
	r.app.Listen(config.Conf.CurrentApp.RestfulAddress)
}

func (r *Restful) Test(req *http.Request) (*http.Response, error) {
	res, err := r.app.Test(req)

	return res, err
}

func (r *Restful) Stop() {
	r.app.Shutdown()
}
