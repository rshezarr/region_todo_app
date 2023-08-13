package app

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"todo_list/internal/config"
	"todo_list/internal/handler"
	"todo_list/internal/repository"
	"todo_list/internal/server"
	"todo_list/internal/service"
	"todo_list/pkg/database/mongodb"
)

func Run() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := mongodb.ConnectDB(context.Background())
	if err != nil {
		log.Fatalf("error while connecting database: %v", err)
	}

	repo := repository.NewRepository(db)
	svc := service.NewService(repo)
	hdr := handler.NewHandler(svc)

	srv := server.NewServer(context.Background(), cfg, hdr.InitRoutes())

	// graceful shutdown
	// creating channel for signals
	quit := make(chan os.Signal, 1)

	// accepting signal from user
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)

	// run server in goroutine
	go func() {
		log.Printf("app: Starting server at port %v -> http:// localhost%v\n", cfg.API.Port, cfg.API.Port)
		srv.Start()
	}()

	// "select" makes the server wait for any signal or error
	select {
	case sig := <-quit:
		log.Printf("app: signal accepted: %v\n", sig)
	case err := <-srv.ServerErrNotify():
		log.Printf("app: server closing: %v\n", err)
	}

	// shutting down the server when any error or signal occurs
	if err := srv.Shutdown(); err != nil {
		log.Printf("error while shutting down server: %s\n", err.Error())
	}
}
