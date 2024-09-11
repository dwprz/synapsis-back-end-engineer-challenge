package restful

import (
	"book-service/src/api/restful/handler"
	"book-service/src/api/restful/server"
	"book-service/src/interface/service"
)

/*
 - Function ini hanya untuk inisialiasi RESTful client dan RESTful server
 - Digunakan untuk injek dependensi dalam lingkungan RESTful (client, delivery, handler, middleware, router, server)
*/

func InitServer(bs service.Book) *server.Restful {
	bookHandler := handler.NewBook(bs)
	restfulServer := server.NewRestful(bookHandler)

	return restfulServer
}
