package authorization

import (
	"context"
	"fmt"
	"github.com/Inspirate789/Thermy-backend/internal/domain/entities"
	"github.com/Inspirate789/Thermy-backend/internal/domain/services/storage"
)

type Session struct {
	authData *entities.AuthRequest
	token    uint64
	role     string
	connDB   storage.ConnDB
}

func NewSession() *Session {
	return &Session{}
}

func (s *Session) GetAuthData() *entities.AuthRequest {
	return s.authData
}

func (s *Session) GetToken() uint64 {
	return s.token
}

func (s *Session) GetRole() string {
	return s.role
}

func (s *Session) GetConn() storage.ConnDB {
	return s.connDB
}

func (s *Session) Open(sm storage.StorageManager, request *entities.AuthRequest, ctx context.Context) (uint64, error) {
	conn, role, err := sm.OpenConn(request, ctx)
	if err != nil {
		return 0, fmt.Errorf("cannot open session: %w", err)
	}

	s.authData = request
	s.token, err = entities.NewUser(request).GetHash()
	s.role = role
	s.connDB = conn

	return s.token, err
}

func (s *Session) Close(sm storage.StorageManager) error {
	err := sm.CloseConn(s.connDB)
	if err != nil {
		return fmt.Errorf("cannot close session: %w", err)
	}

	return nil
}
