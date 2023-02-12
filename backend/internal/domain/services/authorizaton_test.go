package services

import (
	"backend/internal/domain/interfaces"
	"context"
	"testing"
)

type testStorage struct {
}

func (*testStorage) Connect(r *interfaces.AuthRequest, ctx context.Context) error {
	return nil
}

func TestAddSession(t *testing.T) {
	r := interfaces.AuthRequest{
		Username: "TestUser",
		Password: "TestPassword",
	}
	as := NewAuthorizationService()
	var s interfaces.Storage = &testStorage{}

	token, err := as.AddSession(&s, &r, context.Background())
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("Get session token: %d\n", token)
}

func TestRemoveExistingSession(t *testing.T) {
	r := interfaces.AuthRequest{
		Username: "TestUser",
		Password: "TestPassword",
	}

	var s interfaces.Storage = &testStorage{}
	as := NewAuthorizationService()

	token, err := as.AddSession(&s, &r, context.Background())
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("Get session token: %d\n", token)

	err = as.RemoveSession(token)
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestRemoveNonExistingSession(t *testing.T) {
	as := NewAuthorizationService()

	err := as.RemoveSession(12345)
	if err == nil {
		t.Fatalf(err.Error())
	}
	t.Log(err)
}
