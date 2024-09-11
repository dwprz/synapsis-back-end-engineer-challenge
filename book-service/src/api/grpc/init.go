package grpc

import (
	"book-service/src/api/grpc/client"
	"book-service/src/api/grpc/delivery"
	"book-service/src/api/grpc/handler"
	"book-service/src/api/grpc/interceptor"
	"book-service/src/api/grpc/server"
	"book-service/src/infrastructure/config"
	"book-service/src/interface/service"
)

/*
 - Function ini hanya untuk inisialiasi gRPC client dan gRPC server
 - Digunakan untuk injek dependensi dalam lingkungan gRPC (client, delivery, handler, interceptor, server)
*/

func InitServer(bs service.Book) *server.Grpc {
	grpcAddr := config.Conf.CurrentApp.GrpcAddress

	bookGrpcHandler := handler.NewBookGrpc(bs)
	unaryResponseInterceptor := interceptor.NewUnaryResponse()

	grpcServer := server.NewGrpc(grpcAddr, bookGrpcHandler, unaryResponseInterceptor)
	return grpcServer
}

func InitClient() *client.Grpc {
	bookCategoryGrpcDelivery, bookCategoryGrpcConn := delivery.NewBookCategoryGrpc()
	grpcClient := client.NewGrpc(bookCategoryGrpcDelivery, bookCategoryGrpcConn)

	return grpcClient
}
