package grpc

import (
	"book-category-service/src/api/grpc/client"
	"book-category-service/src/api/grpc/delivery"
	"book-category-service/src/api/grpc/handler"
	"book-category-service/src/api/grpc/interceptor"
	"book-category-service/src/api/grpc/server"
	"book-category-service/src/infrastructure/config"
	"book-category-service/src/interface/service"
)

/*
 - Function ini hanya untuk inisialiasi gRPC client dan gRPC server
 - Digunakan untuk injek dependensi dalam lingkungan gRPC (client, delivery, handler, interceptor, server)
*/

func InitClient() *client.Grpc {
	bookGrpcDelivery, bookGrpcConn := delivery.NewBookGrpc()
	grpcClient := client.NewGrpc(bookGrpcDelivery, bookGrpcConn)

	return grpcClient
}

func InitServer(bs service.BookCategory) *server.Grpc {
	 bookCategoryGrpcHandler := handler.NewBookCategoryGrpc(bs)
	 unaryResponseInterceptor := interceptor.NewUnaryResponse()

	grpcServer := server.NewGrpc(config.Conf.CurrentApp.GrpcAddress, bookCategoryGrpcHandler, unaryResponseInterceptor)
	return grpcServer
}