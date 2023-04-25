package authorization

import (
	"context"
	"github.com/Inspirate789/Thermy-backend/internal/domain/entities"
	"github.com/Inspirate789/Thermy-backend/internal/domain/services/storage"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/mock"
	"io"
	"reflect"
	"sync"
	"testing"
)

func TestNewAuthService(t *testing.T) {
	mockLogger := log.New()
	mockLogger.SetOutput(io.Discard)

	tests := []struct {
		name string
		arg  *log.Logger
		want *AuthService
	}{
		{
			name: "Simple positive test",
			arg:  mockLogger,
			want: &AuthService{
				mx:       sync.RWMutex{},
				sessions: make(map[uint64]*Session),
				logger:   mockLogger,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewAuthService(tt.arg)
			if !reflect.DeepEqual(got.logger, tt.want.logger) ||
				!reflect.DeepEqual(got.sessions, tt.want.sessions) {
				t.Errorf("NewAuthService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthService_AddSession(t *testing.T) {
	mockLogger := log.New()
	mockLogger.SetOutput(io.Discard)

	mockSM := new(storage.MockStorageManager)
	mockSM.On("OpenConn", mock.Anything, mock.Anything).Return(nil, "admin", nil)
	mockSM.On("CloseConn", mock.Anything).Return(nil)

	type args struct {
		sm      storage.StorageManager
		request *entities.AuthRequest
		ctx     context.Context
	}
	tests := []struct {
		name    string
		as      *AuthService
		args    args
		want    uint64
		wantErr bool
	}{
		{
			name: "Simple positive test",
			as: &AuthService{
				mx:       sync.RWMutex{},
				sessions: make(map[uint64]*Session),
				logger:   mockLogger,
			},
			args: args{
				sm: mockSM,
				request: &entities.AuthRequest{
					Username: "initial_admin",
					Password: "12345",
				},
				ctx: context.Background(),
			},
			want:    10063865700249539947,
			wantErr: false,
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.as.AddSession(tt.args.sm, tt.args.request, tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddSession() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AddSession() got = %v, want %v", got, tt.want)
			}
		})

		mockSM.AssertNumberOfCalls(t, "OpenConn", i+1)
	}
}

func TestAuthService_RemoveSession(t *testing.T) {
	mockLogger := log.New()
	mockLogger.SetOutput(io.Discard)

	mockSM := new(storage.MockStorageManager)
	mockSM.On("OpenConn", mock.Anything, mock.Anything).Return(nil, "admin", nil)
	mockSM.On("CloseConn", mock.Anything).Return(nil)

	type args struct {
		sm      storage.StorageManager
		request *entities.AuthRequest
		token   uint64
		ctx     context.Context
	}
	tests := []struct {
		name    string
		as      *AuthService
		args    args
		wantErr bool
	}{
		{
			name: "Simple positive test",
			as: &AuthService{
				mx:       sync.RWMutex{},
				sessions: make(map[uint64]*Session),
				logger:   mockLogger,
			},
			args: args{
				sm: mockSM,
				request: &entities.AuthRequest{
					Username: "initial_admin",
					Password: "12345",
				},
				token: 10063865700249539947,
				ctx:   context.Background(),
			},
			wantErr: false,
		},
		{
			name: "Simple negative test",
			as: &AuthService{
				mx:       sync.RWMutex{},
				sessions: make(map[uint64]*Session),
				logger:   mockLogger,
			},
			args: args{
				sm: mockSM,
				request: &entities.AuthRequest{
					Username: "initial_admin",
					Password: "12345",
				},
				token: 1,
				ctx:   context.Background(),
			},
			wantErr: true,
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.as.AddSession(tt.args.sm, tt.args.request, tt.args.ctx)
			if err != nil {
				t.Error(err)
			}
			err = tt.as.RemoveSession(tt.args.sm, tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("RemoveSession() error = %v, wantErr %v", err, tt.wantErr)
			}
		})

		mockSM.AssertNumberOfCalls(t, "OpenConn", i+1)

		if !tt.wantErr {
			mockSM.AssertNumberOfCalls(t, "CloseConn", i+1)
		}
	}
}

func TestAuthService_GetSessionConn(t *testing.T) {
	mockLogger := log.New()
	mockLogger.SetOutput(io.Discard)

	mockSM := new(storage.MockStorageManager)
	mockSM.On("OpenConn", mock.Anything, mock.Anything).Return(nil, "admin", nil)
	mockSM.On("CloseConn", mock.Anything).Return(nil)

	type args struct {
		sm      storage.StorageManager
		request *entities.AuthRequest
		token   uint64
		ctx     context.Context
	}
	tests := []struct {
		name    string
		as      *AuthService
		args    args
		wantErr bool
	}{
		{
			name: "Simple positive test",
			as: &AuthService{
				mx:       sync.RWMutex{},
				sessions: make(map[uint64]*Session),
				logger:   mockLogger,
			},
			args: args{
				sm: mockSM,
				request: &entities.AuthRequest{
					Username: "initial_admin",
					Password: "12345",
				},
				token: 10063865700249539947,
				ctx:   context.Background(),
			},
			wantErr: false,
		},
		{
			name: "Simple negative test",
			as: &AuthService{
				mx:       sync.RWMutex{},
				sessions: make(map[uint64]*Session),
				logger:   mockLogger,
			},
			args: args{
				sm: mockSM,
				request: &entities.AuthRequest{
					Username: "initial_admin",
					Password: "12345",
				},
				token: 1,
				ctx:   context.Background(),
			},
			wantErr: true,
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.as.AddSession(tt.args.sm, tt.args.request, tt.args.ctx)
			if err != nil {
				t.Error(err)
			}
			_, err = tt.as.GetSessionConn(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSessionConn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("GetSessionConn() got = %v, want %v", got, tt.want)
			//}
		})

		mockSM.AssertNumberOfCalls(t, "OpenConn", i+1)
	}
}

func TestAuthService_GetSessionRole(t *testing.T) {
	mockLogger := log.New()
	mockLogger.SetOutput(io.Discard)

	mockSM := new(storage.MockStorageManager)
	mockSM.On("OpenConn", mock.Anything, mock.Anything).Return(nil, "admin", nil)
	mockSM.On("CloseConn", mock.Anything).Return(nil)

	type args struct {
		sm      storage.StorageManager
		request *entities.AuthRequest
		token   uint64
		ctx     context.Context
	}
	tests := []struct {
		name    string
		as      *AuthService
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Simple positive test",
			as: &AuthService{
				mx:       sync.RWMutex{},
				sessions: make(map[uint64]*Session),
				logger:   mockLogger,
			},
			args: args{
				sm: mockSM,
				request: &entities.AuthRequest{
					Username: "initial_admin",
					Password: "12345",
				},
				token: 10063865700249539947,
				ctx:   context.Background(),
			},
			want:    "admin",
			wantErr: false,
		},
		{
			name: "Simple negative test",
			as: &AuthService{
				mx:       sync.RWMutex{},
				sessions: make(map[uint64]*Session),
				logger:   mockLogger,
			},
			args: args{
				sm: mockSM,
				request: &entities.AuthRequest{
					Username: "initial_admin",
					Password: "12345",
				},
				token: 1,
				ctx:   context.Background(),
			},
			want:    "",
			wantErr: true,
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.as.AddSession(tt.args.sm, tt.args.request, tt.args.ctx)
			if err != nil {
				t.Error(err)
			}
			got, err := tt.as.GetSessionRole(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSessionRole() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetSessionRole() got = %v, want %v", got, tt.want)
			}
		})

		mockSM.AssertNumberOfCalls(t, "OpenConn", i+1)
	}
}
