package restful

import (
	"book-category-service/src/api/restful/handler"
	"book-category-service/src/api/restful/server"
	"book-category-service/src/interface/service"
)

/*
 - Function ini hanya untuk inisialiasi RESTful client dan RESTful server
 - Digunakan untuk injek dependensi dalam lingkungan RESTful (client, delivery, handler, middleware, router, server)
*/

func InitServer(bs service.BookCategory) *server.Restful {
	bookCategoryHandler := handler.NewBookCategory(bs)
	restfulServer := server.NewRestful(bookCategoryHandler)

	return restfulServer
}
