package influx_writer

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"
)

func TestInfluxWriter_Write(t *testing.T) {
	w := NewInfluxWriter()
	err := w.Open()
	if err != nil {
		t.Error(err)
		return
	}
	defer w.Close()

	tests := []struct {
		name    string
		arg     []byte
		want    int
		wantErr bool
	}{
		{
			name:    "Simple positive test",
			arg:     []byte("Hello world!"),
			want:    len([]byte("Hello world!")),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := w.Write(tt.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Write() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMain(m *testing.M) {
	err := godotenv.Load("tests.env")
	if err != nil {
		log.Fatal("File tests.env not found")
	}
	code := m.Run()
	// shutdown()
	os.Exit(code)
}
