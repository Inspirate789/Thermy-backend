package authorization

import (
	"context"
	"fmt"
	"github.com/Inspirate789/Thermy-backend/internal/domain/services/storage"
	"github.com/Inspirate789/Thermy-backend/pkg/logger"
	"os"
	"testing"
)

type testStorage struct {
}

func (*testStorage) Connect(r *storage.AuthRequest, ctx context.Context) (storage.ConnDB, string, error) {
	return nil, "TestRole", nil
}

type testLogger struct{}

func (testLogger) Open(serviceName string) error {
	return nil
}

func (testLogger) Print(record logger.LogRecord) {
	fmt.Println(record)
}

func (testLogger) Close() error {
	return nil
}

func TestAddSession(t *testing.T) {
	r := storage.AuthRequest{
		Username: "TestUser",
		Password: "TestPassword",
	}
	s := &testStorage{}
	as := NewAuthorizationService(testLogger{})
	ss := storage.NewStorageService() // TODO

	token, err := as.AddSession(s, &r, context.Background())
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("Get session token: %d\n", token)
}

func TestRemoveExistingSession(t *testing.T) {
	r := storage.AuthRequest{
		Username: "TestUser",
		Password: "TestPassword",
	}

	var s storage.Storage = &testStorage{}
	as := NewAuthorizationService(testLogger{})
	err := as.Open()
	if err != nil {
		t.Fatalf(err.Error())
	}

	token, err := as.AddSession(s, &r, context.Background())
	if err != nil {
		t.Fatalf(err.Error())
	}
	// t.Logf("Get session token: %d\n", token)

	err = as.RemoveSession(token)
	if err != nil {
		t.Fatalf(err.Error())
	}

	err = as.Close()
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestRemoveNonExistingSession(t *testing.T) {
	as := NewAuthorizationService(testLogger{})
	err := as.Open()
	if err != nil {
		t.Fatalf(err.Error())
	}

	err = as.RemoveSession(12345)
	if err == nil {
		t.Fatalf(err.Error())
	}
	t.Log(err)

	err = as.Close()
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestGetSessionRole(t *testing.T) {
	r := storage.AuthRequest{
		Username: "TestUser",
		Password: "TestPassword",
	}

	var s storage.Storage = &testStorage{}
	as := NewAuthorizationService(testLogger{})
	err := as.Open()
	if err != nil {
		t.Fatalf(err.Error())
	}

	token, err := as.AddSession(s, &r, context.Background())
	if err != nil {
		t.Fatalf(err.Error())
	}
	// t.Logf("Get session token: %d\n", token)

	role, err := as.GetSessionRole(token)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if role != "TestRole" {
		t.Fatalf(err.Error())
	}

	err = as.Close()
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestMain(m *testing.M) {
	// ...
	code := m.Run()
	// shutdown()
	os.Exit(code)
}
