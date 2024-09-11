package main

import (
	"book-category-service/src/api/grpc"
	"book-category-service/src/api/restful"
	"book-category-service/src/infrastructure/database"
	"book-category-service/src/repository"
	"book-category-service/src/service"
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

	bookCategoryRepo := repository.NewBookCategory(postgresDB)
	bookCategoryAggr := repository.NewBookCategoryAggregation(postgresDB)

	bookCategoryService := service.NewBookCategory(bookCategoryRepo, bookCategoryAggr, grpcClient)
	
	restfulService := restful.InitServer(bookCategoryService)
	defer restfulService.Stop()

	go restfulService.Run()

	grpcServer := grpc.InitServer(bookCategoryService)
	defer grpcServer.Stop()

	go grpcServer.Run()

	<-ctx.Done()
}
