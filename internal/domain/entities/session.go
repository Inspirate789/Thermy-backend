package entities

import (
	"context"
	"fmt"
	"github.com/Inspirate789/Thermy-backend/internal/domain/interfaces"
	"github.com/Inspirate789/Thermy-backend/internal/domain/services"
	"hash/fnv"
)

type Session struct {
	authData *interfaces.AuthRequest
	token    uint64
	role     string
	connDB   interfaces.ConnDB
}

func (s *Session) GetAuthData() *interfaces.AuthRequest {
	return s.authData
}

func (s *Session) GetToken() uint64 {
	return s.token
}

func (s *Session) GetRole() string {
	return s.role
}

func (s *Session) GetConn() interfaces.ConnDB {
	return s.connDB
}

func generateToken(request *interfaces.AuthRequest) (uint64, error) {
	h := fnv.New64a()

	_, err := h.Write([]byte(request.Username))
	if err != nil {
		return 0, err
	}

	return h.Sum64(), err
}

func (s *Session) Open(ss *services.StorageService, request *interfaces.AuthRequest, ctx context.Context) (uint64, error) {
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

func (s *Session) Close(ss *services.StorageService) error {
	err := ss.CloseConn(s.connDB)
	if err != nil {
		return fmt.Errorf("cannot close session: %w", err)
	}

	return nil
}
