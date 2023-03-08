package storage

import (
	"context"
	"github.com/stretchr/testify/mock"
)

type MockStorageManager struct {
	StorageManager
	mock.Mock
}

func (m *MockStorageManager) OpenConn(request *AuthRequest, ctx context.Context) (ConnDB, string, error) {
	args := m.Called(request, ctx)
	return args.Get(0), "admin", args.Error(2)
}

func (m *MockStorageManager) CloseConn(conn ConnDB) error {
	args := m.Called(conn)
	return args.Error(0)
}
