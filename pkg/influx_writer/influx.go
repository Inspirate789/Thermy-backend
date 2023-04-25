package influx_writer

import (
	"context"
	"errors"
	"os"
	"sync"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/influxdata/influxdb-client-go/v2/domain"
)

type InfluxWriter struct {
	client   influxdb2.Client
	mx       sync.Mutex
	writeAPI api.WriteAPI
}

var (
	influxWriter     *InfluxWriter // Singleton
	influxWriterOnce sync.Once
)

func NewInfluxWriter() *InfluxWriter {
	influxWriterOnce.Do(func() {
		influxWriter = &InfluxWriter{}
	})

	return influxWriter
}

//func NewInfluxWriter() *InfluxWriter {
//	return &InfluxWriter{}
//}

func (il *InfluxWriter) Open() error {
	dbToken, exists := os.LookupEnv("INFLUXDB_TOKEN")
	if !exists {
		return errors.New("INFLUXDB_TOKEN must be set")
	}

	dbURL, exists := os.LookupEnv("INFLUXDB_URL")
	if !exists {
		return errors.New("INFLUXDB_URL must be set")
	}

	dbOrg, exists := os.LookupEnv("INFLUXDB_ORG")
	if !exists {
		return errors.New("INFLUXDB_ORG must be set")
	}

	dbBucket, exists := os.LookupEnv("INFLUXDB_BACKEND_BUCKET_NAME")
	if !exists {
		return errors.New("INFLUXDB_BACKEND_BUCKET_NAME must be set")
	}

	client := influxdb2.NewClient(dbURL, dbToken)

	health, err := client.Health(context.Background()) // validate client connection health
	if (err != nil) && health.Status == domain.HealthCheckStatusPass {
		client.Close()
		return errors.New("connectToInfluxDB() error: database not healthy")
	}

	il.client = client
	il.writeAPI = il.client.WriteAPI(dbOrg, dbBucket) // Get non-blocking write client

	return nil
}

func (il *InfluxWriter) Write(p []byte) (int, error) {
	point := influxdb2.NewPointWithMeasurement("Messages").
		AddField("message", p).
		SetTime(time.Now())

	il.mx.Lock()
	il.writeAPI.WritePoint(point)
	// Flush writes
	il.writeAPI.Flush()
	il.mx.Unlock()

	return len(p), nil
}

func (il *InfluxWriter) Close() {
	il.client.Close()
}
