package logger

import (
	"context"
	"errors"
	"fmt"
	"os"
	"sync"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/influxdata/influxdb-client-go/v2/domain"
)

type InfluxLogger struct {
	client   influxdb2.Client
	mx       sync.Mutex
	writeAPI api.WriteAPI
}

var (
	influxLogger     *InfluxLogger // Singleton
	influxLoggerOnce sync.Once
)

func NewInfluxLogger() *InfluxLogger {
	influxLoggerOnce.Do(func() {
		influxLogger = &InfluxLogger{}
	})

	return influxLogger
}

//func NewInfluxLogger() *InfluxLogger {
//	return &InfluxLogger{}
//}

func (il *InfluxLogger) connectToInfluxDB() (influxdb2.Client, error) {
	dbToken, exists := os.LookupEnv("INFLUXDB_TOKEN")
	if !exists {
		return nil, errors.New("INFLUXDB_TOKEN must be set")
	}

	dbURL, exists := os.LookupEnv("INFLUXDB_URL")
	if !exists {
		return nil, errors.New("INFLUXDB_URL must be set")
	}

	// TODO: Enabling SSL/TLS encryption
	client := influxdb2.NewClient(dbURL, dbToken)

	return client, nil
}

func (il *InfluxLogger) Open(serviceName string) error {
	client, err := il.connectToInfluxDB()
	if err != nil {
		return err
	}

	// validate client connection health
	health, err := client.Health(context.Background())
	if (err != nil) && health.Status == domain.HealthCheckStatusPass {
		return errors.New("connectToInfluxDB() error. database not healthy")
	}

	il.client = client

	dbORG := os.Getenv("INFLUXDB_ORG")
	if dbORG == "" {
		return errors.New("INFLUXDB_ORG must be set")
	}

	// Get non-blocking write client
	il.writeAPI = il.client.WriteAPI(dbORG, serviceName)

	return nil
}

func (il *InfluxLogger) Print(r LogRecord) {
	// Add data point
	p := influxdb2.NewPointWithMeasurement(fmt.Sprintf("%s", r.Type)).
		AddField("name", r.Name).
		AddField("type", r.Type).
		AddField("message", r.Msg).
		SetTime(time.Now())

	il.mx.Lock()
	il.writeAPI.WritePoint(p)
	// Flush writes
	il.writeAPI.Flush() // TODO: add flush logic (flush with period?)
	il.mx.Unlock()
}

func (il *InfluxLogger) Close() {
	il.client.Close()
}
