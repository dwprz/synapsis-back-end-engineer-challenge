package main

import (
	"book-service/src/api/grpc"
	"book-service/src/api/restful"
	"book-service/src/infrastructure/database"
	"book-service/src/repository"
	"book-service/src/service"
	"context"
	"os"
	"os/signal"
	"syscall"
)

func handleCloseApp(cancel context.CancelFunc) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		cancel()
	}()
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	handleCloseApp(cancel)

	postgresDB := database.NewPostgres()
	defer postgresDB.Close()

	grpcClient := grpc.InitClient()
	defer grpcClient.Close()

	bookRepo := repository.NewBook(postgresDB, grpcClient)
	bookAggrRepo := repository.NewBookAggregation(postgresDB)
	popularTitleKeyRepo := repository.NewPopularTitleKey(postgresDB)

	bookService := service.NewBook(bookRepo, bookAggrRepo, popularTitleKeyRepo)

	grpcServer := grpc.InitServer(bookService)
	defer grpcServer.Stop()

	go grpcServer.Run()

	restfulServer := restful.InitServer(bookService)
	defer restfulServer.Stop()

	go restfulServer.Run()

	<-ctx.Done()
}
