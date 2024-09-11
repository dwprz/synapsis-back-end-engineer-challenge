package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"user-service/src/api/restful"
	"user-service/src/cache"
	"user-service/src/infrastructure/database"
	"user-service/src/repository"
	"user-service/src/service"
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

	redisDB := database.NewRedisCluster()
	defer redisDB.Close()

	userCache := cache.NewUser(redisDB)
	userRepo := repository.NewUser(postgresDB, userCache)

	authService := service.NewAuth(userRepo, userCache)
	userService := service.NewUser(userRepo, userCache)

	restfulServer := restful.InitServer(authService, userService)
	defer restfulServer.Stop()

	go restfulServer.Run()

	<-ctx.Done()
}
