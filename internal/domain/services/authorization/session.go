package authorization

import (
	"context"
	"github.com/Inspirate789/Thermy-backend/internal/domain/entities"
	"github.com/Inspirate789/Thermy-backend/internal/domain/services/storage"
)

type session struct {
	authData *entities.AuthRequest
	token    uint64
	role     string
	connDB   storage.ConnDB
}

func newSession() *session {
	return &session{}
}

func (s *session) GetAuthData() *entities.AuthRequest {
	return s.authData
}

func (s *session) GetToken() uint64 {
	return s.token
}

func (s *session) GetRole() string {
	return s.role
}

func (s *session) GetConn() storage.ConnDB {
	return s.connDB
}

func (s *session) Open(sm storage.StorageManager, request *entities.AuthRequest, ctx context.Context) (uint64, error) {
	conn, role, err := sm.OpenConn(request, ctx)
	if err != nil {
		return 0, err
	}

	s.authData = request
	s.token, err = entities.NewUser(request).GetHash()
	s.role = role
	s.connDB = conn

	return s.token, err
}

func (s *session) Close(sm storage.StorageManager) error {
	err := sm.CloseConn(s.connDB)
	if err != nil {
		return err
	}

	return nil
}
