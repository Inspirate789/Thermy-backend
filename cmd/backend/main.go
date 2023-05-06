package main

import (
	"flag"
	"fmt"
	"github.com/Inspirate789/Thermy-backend/internal/adapters/server"
	"github.com/Inspirate789/Thermy-backend/internal/adapters/storage/postgres_storage"
	"github.com/Inspirate789/Thermy-backend/internal/domain/services/authorization"
	"github.com/Inspirate789/Thermy-backend/internal/domain/services/storage"
	influx "github.com/Inspirate789/Thermy-backend/pkg/influx_writer"
	_ "github.com/Inspirate789/Thermy-backend/swagger"
	runtime "github.com/banzaicloud/logrus-runtime-formatter"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

var configFilename string
var startConfig []byte

func init() {
	flag.StringVar(&configFilename, "env", "backend.env", ".env file for backend")
	flag.Parse()

	var err error
	startConfig, err = os.ReadFile(configFilename)
	if err != nil {
		log.Fatalf("File %s not readed: %v", configFilename, err)
	}
	err = godotenv.Load(configFilename) // TODO: read filename from flag
	if err != nil {
		log.Fatalf("File %s not loaded: %v", configFilename, err)
	}

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

func exitServer(logger *log.Logger, srv *server.Server) {
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Shutdown Server ...")

	err := srv.Stop()
	if err != nil {
		logger.Errorf("Server Shutdown: %v", err)
	}
	logger.Info("Server exited")
}

func NewLogger(w io.Writer) *log.Logger {
	level, err := log.ParseLevel(os.Getenv("BACKEND_LOGLEVEL"))
	if err != nil {
		log.Fatal(err)
	}

	logger := log.New()
	logger.SetOutput(w)
	logger.SetLevel(level)
	formatter := runtime.Formatter{
		ChildFormatter: &log.TextFormatter{
			FullTimestamp: true,
		},
		Package: true,
		File:    true,
		Line:    true,
	}
	formatter.Line = true
	logger.SetFormatter(&formatter)

	return logger
}

//	@title			Thermy API
//	@version		1.0
//	@description	This is a Thermy backend API.

//	@contact.name	API Support
//	@contact.email	andreysapozhkov535@gmail.com

//	@license.name	MIT
//	@license.url	https://mit-license.org/

//	@host		localhost:8080
//	@BasePath	/api/v1
//	@Schemes	http
func main() {
	w := influx.NewInfluxWriter()
	err := w.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer w.Close()
	logger := NewLogger(w)

	authService := authorization.NewAuthService(logger)
	storageService := storage.NewStorageService(postgres_storage.NewPostgresStorage(), logger)

	port, err := strconv.Atoi(os.Getenv("BACKEND_PORT"))
	if err != nil {
		logger.Fatal(err)
	}

	srv := server.NewServer(port, authService, storageService, logger)

	go func() {
		err = srv.Start()
		if err != nil && err != http.ErrServerClosed {
			logger.Error(fmt.Sprintf("listen: %s\n", err))
		}
	}()
	logger.Infof("Server started at port %s with configuration: \n%s", os.Getenv("BACKEND_PORT"), string(startConfig))

	exitServer(logger, srv)
}
