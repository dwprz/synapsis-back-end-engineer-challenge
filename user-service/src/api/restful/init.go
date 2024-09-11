package restful

import (
	"user-service/src/api/restful/handler"
	"user-service/src/api/restful/server"
	"user-service/src/interface/service"
)

/*
 - Function ini hanya untuk inisialiasi RESTful client dan RESTful server
 - Digunakan untuk injek dependensi dalam lingkungan RESTful (client, delivery, handler, middleware, router, server)
*/

func InitServer(as service.Auth, us service.User) *server.Restful {
	authHandler := handler.NewAuth(as)
	userHandler := handler.NewUser(us)

	restfulServer := server.NewRestful(authHandler, userHandler)
	return restfulServer
}
