package main

import (
	"fmt"
	"github.com/Inspirate789/Thermy-backend/internal/adapters/server"
	"github.com/Inspirate789/Thermy-backend/internal/adapters/storage/postgres_storage"
	"github.com/Inspirate789/Thermy-backend/internal/domain/services"
	"github.com/Inspirate789/Thermy-backend/pkg/logger"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	// TODO: read .env file and set environment variable

	if err := godotenv.Load("main.env"); err != nil {
		log.Fatal("No .env file found")
	}
}

func main() { // TODO: decompose main into initServer, startServer, stopServer?
	mainLog := logger.NewInfluxLogger()
	err := mainLog.Open("Backend")
	if err != nil {
		log.Fatal(err)
	}
	defer mainLog.Close()

	authService := services.NewAuthorizationService(mainLog)
	storageService := services.NewStorageService(postgres_storage.NewPostgresStorage(), mainLog)
	srv := server.NewServer(8080, &authService, &storageService, mainLog)

	go func() {
		err = srv.Start()
		if err != nil && err != http.ErrServerClosed {
			mainLog.Print(logger.LogRecord{
				Name: "Main",
				Type: logger.Error,
				Msg:  fmt.Sprintf("listen: %s\n", err),
			})
		}
	}()

	mainLog.Print(logger.LogRecord{
		Name: "Main",
		Type: logger.Debug,
		Msg:  "Server started",
	})

	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	mainLog.Print(logger.LogRecord{
		Name: "Main",
		Type: logger.Debug,
		Msg:  "Shutdown Server ...",
	})
	err = srv.Stop()
	if err != nil {
		mainLog.Print(logger.LogRecord{
			Name: "Main",
			Type: logger.Error,
			Msg:  fmt.Sprintf("Server Shutdown: %v", err),
		})
	}
	mainLog.Print(logger.LogRecord{
		Name: "Main",
		Type: logger.Debug,
		Msg:  "Server exited",
	})
}
