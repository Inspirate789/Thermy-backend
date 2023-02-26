package authorization

import (
	"context"
	"fmt"
	"github.com/Inspirate789/Thermy-backend/internal/domain/services/storage"
	"hash/fnv"
)

type Session struct {
	authData *storage.AuthRequest
	token    uint64
	role     string
	connDB   storage.ConnDB
}

func (s *Session) GetAuthData() *storage.AuthRequest {
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

func generateToken(request *storage.AuthRequest) (uint64, error) {
	h := fnv.New64a()

	_, err := h.Write([]byte(request.Username))
	if err != nil {
		return 0, err
	}

	return h.Sum64(), err
}

func (s *Session) Open(ss *storage.StorageService, request *storage.AuthRequest, ctx context.Context) (uint64, error) {
	conn, role, err := ss.OpenConn(request, ctx)
	if err != nil {
		return 0, fmt.Errorf("cannot open session: %w", err)
	}

	s.authData = request
	s.token, err = generateToken(request)
	s.role = role
	s.connDB = conn

	return s.token, err
}

func (s *Session) Close(ss *storage.StorageService) error {
	err := ss.CloseConn(s.connDB)
	if err != nil {
		return fmt.Errorf("cannot close session: %w", err)
	}

	return nil
}
