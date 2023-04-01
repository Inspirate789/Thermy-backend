package authorization

import (
	"context"
	"fmt"
	"github.com/Inspirate789/Thermy-backend/internal/domain/entities"
	"github.com/Inspirate789/Thermy-backend/internal/domain/errors"
	"github.com/Inspirate789/Thermy-backend/internal/domain/services/storage"
	"github.com/Inspirate789/Thermy-backend/pkg/logger"
	"sync"
)

type AuthService struct {
	mx       sync.RWMutex
	sessions map[uint64]*Session
	log      logger.Logger
}

func NewAuthService(log logger.Logger) *AuthService {
	return &AuthService{
		mx:       sync.RWMutex{},
		sessions: make(map[uint64]*Session),
		log:      log,
	}
}

func (as *AuthService) AddSession(sm storage.StorageManager, request *entities.AuthRequest, ctx context.Context) (uint64, error) {
	session := NewSession()
	token, err := session.Open(sm, request, ctx)
	if err != nil {
		return 0, err
	}

	as.mx.Lock()
	as.sessions[token] = session
	as.mx.Unlock()

	as.log.Print(logger.LogRecord{
		Name: "AuthService",
		Type: logger.Debug,
		Msg:  fmt.Sprintf("Add session with token %d, role %s", session.GetToken(), session.GetRole()),
	})

	return token, nil
}

func (as *AuthService) RemoveSession(sm storage.StorageManager, token uint64) error {
	as.mx.RLock()
	session, ok := as.sessions[token]
	as.mx.RUnlock()
	if !ok {
		return errors.ErrRemoveSession
	}

	err := session.Close(sm)
	if err != nil {
		return err // TODO: wrap errors on every layer
	}

	as.log.Print(logger.LogRecord{
		Name: "AuthService",
		Type: logger.Debug,
		Msg:  fmt.Sprintf("Remove session with token %d, role %s", session.GetToken(), session.GetRole()),
	})

	as.mx.Lock()
	delete(as.sessions, token)
	as.mx.Unlock()

	return nil
}

func (as *AuthService) GetSessionRole(token uint64) (string, error) {
	as.mx.RLock()
	session, ok := as.sessions[token]
	as.mx.RUnlock()
	if !ok {
		return "", errors.ErrGetSession
	}

	return session.GetRole(), nil
}

func (as *AuthService) GetSessionConn(token uint64) (storage.ConnDB, error) {
	as.mx.RLock()
	session, ok := as.sessions[token]
	as.mx.RUnlock()
	if !ok {
		return "", errors.ErrGetSession
	}

	return session.GetConn(), nil
}
