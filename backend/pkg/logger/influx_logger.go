package logger

import (
	"context"
	"errors"
	"fmt"
	"os"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/influxdata/influxdb-client-go/v2/domain"
)

type InfluxLogger struct {
	client   influxdb2.Client
	writeAPI api.WriteAPI
}

func NewInfluxLogger() *InfluxLogger {
	return &InfluxLogger{}
}

func (il *InfluxLogger) connectToInfluxDB() (influxdb2.Client, error) {
	dbToken := os.Getenv("INFLUXDB_TOKEN")
	if dbToken == "" {
		return nil, errors.New("INFLUXDB_TOKEN must be set")
	}

	dbURL := os.Getenv("INFLUXDB_URL")
	if dbURL == "" {
		return nil, errors.New("INFLUXDB_URL must be set")
	}

	// TODO: Enabling SSL/TLS encryption
	client := influxdb2.NewClient(dbURL, dbToken)

	return client, nil
}

func (il *InfluxLogger) Open() error {
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

	// get non-blocking write client
	il.writeAPI = il.client.WriteAPI("thermy-org", "thermy-bucket")

	return nil
}

func (il *InfluxLogger) Print(r LogRecord) {
	// write line protocol
	il.writeAPI.WriteRecord(fmt.Sprintf("%s,type=%s,message=%s", r.Name, r.Type, r.Msg))
	// Flush writes
	il.writeAPI.Flush() // TODO: add flush logic (flush with period)
}

func (il *InfluxLogger) Close() error {
	il.client.Close()

	return nil
}
