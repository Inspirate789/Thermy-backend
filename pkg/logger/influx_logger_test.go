package logger

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"
)

//func TestInfluxLogger_connectToInfluxDB(t *testing.T) {
//	il := NewInfluxLogger()
//	client, err := il.connectToInfluxDB()
//	if err != nil {
//		t.Fatalf("connectToInfluxDB() error = %v", err)
//	}
//	if client == nil {
//		t.Fatal("connectToInfluxDB() give empty client")
//	}
//}

func TestInfluxLogger_Open(t *testing.T) {
	il := NewInfluxLogger()
	err := il.Open("Test")
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}

	il.Close()
}

func TestInfluxLogger_Print(t *testing.T) {
	il := NewInfluxLogger()
	err := il.Open("Test")
	if err != nil {
		t.Fatalf("Print() error = %v", err)
	}
	tests := []struct {
		name   string
		logger *InfluxLogger
		arg    LogRecord
	}{
		{
			name:   "Debug",
			logger: il,
			arg: LogRecord{
				Name: "TestService1",
				Type: Debug,
				Msg:  "Debug Message 12345!",
			},
		},
		{
			name:   "Warning",
			logger: il,
			arg: LogRecord{
				Name: "TestService2",
				Type: Warning,
				Msg:  "Warning Message 12345!",
			},
		},
		{
			name:   "Error",
			logger: il,
			arg: LogRecord{
				Name: "TestService3",
				Type: Error,
				Msg:  "Error Message 12345!",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ttlogger := tt.logger
			ttlogger.Print(tt.arg)
		})
	}

	il.Close()
}

func TestMain(m *testing.M) {
	if err := godotenv.Load("tests.env"); err != nil {
		log.Fatal("No .env file found")
	}
	code := m.Run()
	// shutdown()
	os.Exit(code)
}
