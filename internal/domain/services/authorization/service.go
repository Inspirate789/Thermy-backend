package authorization

import (
	"context"
	"github.com/Inspirate789/Thermy-backend/internal/domain/entities"
	"github.com/Inspirate789/Thermy-backend/internal/domain/errors"
	"github.com/Inspirate789/Thermy-backend/internal/domain/services/storage"
	log "github.com/sirupsen/logrus"
	"sync"
)

type AuthService struct {
	mx       sync.RWMutex
	sessions map[uint64]*session
	logger   *log.Logger
}

func NewAuthService(logger *log.Logger) *AuthService {
	return &AuthService{
		mx:       sync.RWMutex{},
		sessions: make(map[uint64]*session),
		logger:   logger,
	}
}

func (as *AuthService) AddSession(sm storage.StorageManager, request *entities.AuthRequest, ctx context.Context) (uint64, error) {
	s := newSession()
	token, err := s.Open(sm, request, ctx)
	if err != nil {
		as.logger.Error(err)
		return 0, errors.ErrOpenSession
	}

	as.mx.Lock()
	as.sessions[token] = s
	as.mx.Unlock()

	as.logger.Infof("AuthService: add session with token %d, role %s", s.GetToken(), s.GetRole())

	return token, nil
}

func (as *AuthService) RemoveSession(sm storage.StorageManager, token uint64) error {
	as.mx.RLock()
	s, ok := as.sessions[token]
	as.mx.RUnlock()
	if !ok {
		as.logger.Error(errors.ErrRemoveSessionByToken)
		return errors.ErrRemoveSessionByToken
	}

	err := s.Close(sm)
	if err != nil {
		as.logger.Error(err)
		return errors.ErrRemoveDatabaseSession
	}

	as.logger.Infof("AuthService: remove session with token %d, role %s", s.GetToken(), s.GetRole())

	as.mx.Lock()
	delete(as.sessions, token)
	as.mx.Unlock()

	return nil
}

func (as *AuthService) GetSessionRole(token uint64) (string, error) {
	as.mx.RLock()
	s, ok := as.sessions[token]
	as.mx.RUnlock()
	if !ok {
		as.logger.Error(errors.ErrGetSession)
		return "", errors.ErrGetSession
	}

	return s.GetRole(), nil
}

func (as *AuthService) GetSessionConn(token uint64) (storage.ConnDB, error) {
	as.mx.RLock()
	s, ok := as.sessions[token]
	as.mx.RUnlock()
	if !ok {
		as.logger.Error(errors.ErrGetSession)
		return "", errors.ErrGetSession
	}

	return s.GetConn(), nil
}

func (as *AuthService) SessionExist(token uint64) bool {
	as.mx.RLock()
	_, exist := as.sessions[token]
	as.mx.RUnlock()

	return exist
}
