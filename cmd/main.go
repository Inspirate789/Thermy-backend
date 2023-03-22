package main

import (
	"fmt"
	"github.com/Inspirate789/Thermy-backend/internal/adapters/server"
	"github.com/Inspirate789/Thermy-backend/internal/adapters/storage/postgres_storage"
	"github.com/Inspirate789/Thermy-backend/internal/domain/services/authorization"
	"github.com/Inspirate789/Thermy-backend/internal/domain/services/storage"
	"github.com/Inspirate789/Thermy-backend/pkg/logger"
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func init() {
	err := godotenv.Load("backend.env") // TODO: read filename from flag
	if err != nil {
		panic("File backend.env not found")
	}

	//err = os.Remove("backend.env")
	//if err != nil {
	//	panic(err)
	//}

	initTimeStr := os.Getenv("BACKEND_INIT_SLEEP_TIME")
	if initTimeStr == "" {
		panic("BACKEND_INIT_SLEEP_TIME must be set")
	}
	initTime, err := strconv.Atoi(initTimeStr)
	if err != nil {
		panic(err)
	}
	time.Sleep(time.Duration(initTime) * time.Second)
}

func exitServer(mainLog logger.Logger, srv *server.Server) {
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

	err := srv.Stop()
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

func main() { // TODO: decompose main into initServer, startServer, stopServer?
	mainLog := logger.NewInfluxLogger()

	logBucketName := os.Getenv("INFLUXDB_BACKEND_BUCKET_NAME")
	if logBucketName == "" {
		panic("INFLUXDB_BACKEND_BUCKET_NAME must be set")
	}

	err := mainLog.Open(logBucketName)
	if err != nil {
		panic(err)
	}
	defer mainLog.Close()

	authService := authorization.NewAuthService(mainLog)
	storageService := storage.NewStorageService(postgres_storage.NewPostgresStorage(), mainLog)

	portStr := os.Getenv("BACKEND_PORT")
	if portStr == "" {
		panic("BACKEND_PORT must be set")
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		panic(err)
	}

	srv := server.NewServer(port, authService, storageService, mainLog)

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

	exitServer(mainLog, srv)
}
