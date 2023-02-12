package services

import (
	"backend/internal/domain/entities" // TODO: replace relational paths by github.com/...
	"backend/internal/domain/errors"
	"backend/internal/domain/interfaces"
	"backend/pkg/logger"
	"context"
	"fmt"
	"hash/fnv"
)

type AuthorizationService struct {
	sessions map[uint64]entities.Session
	log      logger.Logger
}

func NewAuthorizationService() AuthorizationService {
	return AuthorizationService{
		sessions: map[uint64]entities.Session{},
		log:      nil,
	}
}

func (as *AuthorizationService) generateToken(r *interfaces.AuthRequest) uint64 {
	h := fnv.New64a()
	h.Write([]byte(r.Password))

	return h.Sum64()
}

func (as *AuthorizationService) AddSession(db *interfaces.Storage, r *interfaces.AuthRequest, ctx context.Context) (uint64, error) {
	var session entities.Session
	token, err := session.Open(db, r, ctx)
	if err != nil {
		return 0, err
	}
	as.sessions[token] = session

	as.log.Print(logger.LogRecord{
		Name: "AuthorizationService",
		Type: logger.Debug,
		Msg:  fmt.Sprintf("Add session with token %d, role %s", session.Token, session.Role),
	})

	return token, nil
}

func (as *AuthorizationService) RemoveSession(token uint64) error {
	session, ok := as.sessions[token]
	if !ok {
		return errors.ErrRemoveSession
	}

	as.log.Print(logger.LogRecord{
		Name: "AuthorizationService",
		Type: logger.Debug,
		Msg:  fmt.Sprintf("Remove session with token %d, role %s", session.Token, session.Role),
	})

	delete(as.sessions, token)

	return nil
}

func (as *AuthorizationService) GetSessionRole(token uint64) (string, error) {
	session, ok := as.sessions[token]
	if !ok {
		return "", errors.ErrGetSession
	}

	return session.Role, nil
}
