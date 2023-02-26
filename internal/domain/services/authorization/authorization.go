package authorization

import (
	"context"
	"fmt"
	"github.com/Inspirate789/Thermy-backend/internal/domain/errors"
	"github.com/Inspirate789/Thermy-backend/internal/domain/services/storage"
	"github.com/Inspirate789/Thermy-backend/pkg/logger"
	"sync"
)

type AuthorizationService struct {
	mx       sync.RWMutex
	sessions map[uint64]Session
	log      logger.Logger
}

func NewAuthorizationService(log logger.Logger) AuthorizationService {
	return AuthorizationService{
		sessions: make(map[uint64]Session),
		log:      log,
	}
}

func (as *AuthorizationService) AddSession(ss *storage.StorageService, request *storage.AuthRequest, ctx context.Context) (uint64, error) {
	var session Session
	token, err := session.Open(ss, request, ctx)
	if err != nil {
		return 0, err
	}

	as.mx.Lock()
	as.sessions[token] = session
	as.mx.Unlock()

	as.log.Print(logger.LogRecord{
		Name: "AuthorizationService",
		Type: logger.Debug,
		Msg:  fmt.Sprintf("Add session with token %d, role %s", session.GetToken(), session.GetRole()),
	})

	return token, nil
}

func (as *AuthorizationService) RemoveSession(ss *storage.StorageService, token uint64) error {
	as.mx.RLock()
	session, ok := as.sessions[token]
	as.mx.RUnlock()
	if !ok {
		return errors.ErrRemoveSession
	}

	err := session.Close(ss)
	if err != nil {
		return err // TODO: wrap errors on every layer
	}

	as.log.Print(logger.LogRecord{
		Name: "AuthorizationService",
		Type: logger.Debug,
		Msg:  fmt.Sprintf("Remove session with token %d, role %s", session.GetToken(), session.GetRole()),
	})

	as.mx.Lock()
	delete(as.sessions, token)
	as.mx.Unlock()

	return nil
}

func (as *AuthorizationService) GetSessionRole(token uint64) (string, error) {
	as.mx.RLock()
	session, ok := as.sessions[token]
	as.mx.RUnlock()
	if !ok {
		return "", errors.ErrGetSession
	}

	return session.GetRole(), nil
}

func (as *AuthorizationService) GetSessionConn(token uint64) (storage.ConnDB, error) {
	as.mx.RLock()
	session, ok := as.sessions[token]
	as.mx.RUnlock()
	if !ok {
		return "", errors.ErrGetSession
	}

	return session.GetConn(), nil
}
