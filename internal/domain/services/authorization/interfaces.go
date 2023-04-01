package authorization

import (
	"context"
	"github.com/Inspirate789/Thermy-backend/internal/domain/entities"
	"github.com/Inspirate789/Thermy-backend/internal/domain/services/storage"
)

type AuthManager interface {
	AddSession(sm storage.StorageManager, request *entities.AuthRequest, ctx context.Context) (uint64, error)
	RemoveSession(sm storage.StorageManager, token uint64) error
	GetSessionRole(token uint64) (string, error)
	GetSessionConn(token uint64) (storage.ConnDB, error)
}
