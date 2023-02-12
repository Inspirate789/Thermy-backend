package entities

import (
	"backend/internal/domain/interfaces"
	"context"
	"fmt"
	"hash/fnv"
)

type Session struct {
	AuthData *interfaces.AuthRequest
	Token    uint64
	Role     string
}

func (s *Session) generateToken(r *interfaces.AuthRequest) uint64 {
	h := fnv.New64a()
	h.Write([]byte(r.Password))

	return h.Sum64()
}

func (s *Session) Open(db *interfaces.Storage, r *interfaces.AuthRequest, ctx context.Context) (uint64, error) {
	role, err := (*db).Connect(r, ctx)
	if err != nil {
		return 0, fmt.Errorf("Cannot open session: %w", err)
	}

	s.AuthData = r
	s.Token = s.generateToken(r)
	s.Role = role

	return s.Token, nil
}
