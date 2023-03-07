package authorization

import (
	"context"
	"github.com/Inspirate789/Thermy-backend/internal/domain/services/storage"
	"github.com/Inspirate789/Thermy-backend/pkg/logger"
	"github.com/stretchr/testify/mock"
	"reflect"
	"sync"
	"testing"
)

type mockLogger struct {
	mock.Mock
}

func (m *mockLogger) Open(serviceName string) error {
	args := m.Called(serviceName)
	return args.Error(0)
}

func (m *mockLogger) Print(r logger.LogRecord) {
	m.Called(r)
}

func (m *mockLogger) Close() {
	m.Called()
}

type mockStorageManager struct {
	storage.StorageManager
	mock.Mock
}

func (m *mockStorageManager) OpenConn(request *storage.AuthRequest, ctx context.Context) (storage.ConnDB, string, error) {
	args := m.Called(request, ctx)
	return args.Get(0), "admin", args.Error(2)
}

func (m *mockStorageManager) CloseConn(conn storage.ConnDB) error {
	args := m.Called(conn)
	return args.Error(0)
}

func TestNewAuthService(t *testing.T) {
	mockLog := new(mockLogger)
	mockLog.On("Print", mock.Anything).Return()

	tests := []struct {
		name string
		arg  logger.Logger
		want AuthService
	}{
		{
			name: "Simple positive test",
			arg:  mockLog,
			want: AuthService{
				mx:       sync.RWMutex{},
				sessions: make(map[uint64]*Session),
				log:      mockLog,
			},
		},
	}
	for i := range tests {
		t.Run(tests[i].name, func(t *testing.T) {
			got := NewAuthService(tests[i].arg)
			if !reflect.DeepEqual(got.log, tests[i].want.log) ||
				!reflect.DeepEqual(got.sessions, tests[i].want.sessions) {
				t.Errorf("NewAuthService() = %v, want %v", got, tests[i].want)
			}
		})

		mockLog.AssertNumberOfCalls(t, "Open", 0)
		mockLog.AssertNumberOfCalls(t, "Print", 0)
		mockLog.AssertNumberOfCalls(t, "Close", 0)
	}
}

func TestAuthService_AddSession(t *testing.T) {
	mockLog := new(mockLogger)
	mockLog.On("Print", mock.Anything).Return()

	mockSM := new(mockStorageManager)
	mockSM.On("OpenConn", mock.Anything, mock.Anything).Return(nil, "admin", nil)
	mockSM.On("CloseConn", mock.Anything).Return(nil)

	type args struct {
		sm      storage.StorageManager
		request *storage.AuthRequest
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
				log:      mockLog,
			},
			args: args{
				sm: mockSM,
				request: &storage.AuthRequest{
					Username: "initial_admin",
					Password: "12345",
				},
				ctx: context.Background(),
			},
			want:    10063865700249539947,
			wantErr: false,
		},
	}
	for i := range tests {
		t.Run(tests[i].name, func(t *testing.T) {
			got, err := tests[i].as.AddSession(tests[i].args.sm, tests[i].args.request, tests[i].args.ctx)
			if (err != nil) != tests[i].wantErr {
				t.Errorf("AddSession() error = %v, wantErr %v", err, tests[i].wantErr)
				return
			}
			if got != tests[i].want {
				t.Errorf("AddSession() got = %v, want %v", got, tests[i].want)
			}
		})

		mockLog.AssertNumberOfCalls(t, "Open", 0)
		mockLog.AssertNumberOfCalls(t, "Print", i+1)
		mockLog.AssertNumberOfCalls(t, "Close", 0)

		mockSM.AssertNumberOfCalls(t, "OpenConn", i+1)
	}
}

func TestAuthService_RemoveSession(t *testing.T) {
	mockLog := new(mockLogger)
	mockLog.On("Print", mock.Anything).Return()

	mockSM := new(mockStorageManager)
	mockSM.On("OpenConn", mock.Anything, mock.Anything).Return(nil, "admin", nil)
	mockSM.On("CloseConn", mock.Anything).Return(nil)

	type args struct {
		sm      storage.StorageManager
		request *storage.AuthRequest
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
				log:      mockLog,
			},
			args: args{
				sm: mockSM,
				request: &storage.AuthRequest{
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
				log:      mockLog,
			},
			args: args{
				sm: mockSM,
				request: &storage.AuthRequest{
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

		mockLog.AssertNumberOfCalls(t, "Open", 0)
		mockLog.AssertNumberOfCalls(t, "Close", 0)

		mockSM.AssertNumberOfCalls(t, "OpenConn", i+1)
	}
}

func TestAuthService_GetSessionConn(t *testing.T) {
	mockLog := new(mockLogger)
	mockLog.On("Print", mock.Anything).Return()

	mockSM := new(mockStorageManager)
	mockSM.On("OpenConn", mock.Anything, mock.Anything).Return(nil, "admin", nil)
	mockSM.On("CloseConn", mock.Anything).Return(nil)

	type args struct {
		sm      storage.StorageManager
		request *storage.AuthRequest
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
				log:      mockLog,
			},
			args: args{
				sm: mockSM,
				request: &storage.AuthRequest{
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
				log:      mockLog,
			},
			args: args{
				sm: mockSM,
				request: &storage.AuthRequest{
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

		mockLog.AssertNumberOfCalls(t, "Open", 0)
		mockLog.AssertNumberOfCalls(t, "Print", i+1)
		mockLog.AssertNumberOfCalls(t, "Close", 0)

		mockSM.AssertNumberOfCalls(t, "OpenConn", i+1)
	}
}

func TestAuthService_GetSessionRole(t *testing.T) {
	mockLog := new(mockLogger)
	mockLog.On("Print", mock.Anything).Return()

	mockSM := new(mockStorageManager)
	mockSM.On("OpenConn", mock.Anything, mock.Anything).Return(nil, "admin", nil)
	mockSM.On("CloseConn", mock.Anything).Return(nil)

	type args struct {
		sm      storage.StorageManager
		request *storage.AuthRequest
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
				log:      mockLog,
			},
			args: args{
				sm: mockSM,
				request: &storage.AuthRequest{
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
				log:      mockLog,
			},
			args: args{
				sm: mockSM,
				request: &storage.AuthRequest{
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

		mockLog.AssertNumberOfCalls(t, "Open", 0)
		mockLog.AssertNumberOfCalls(t, "Print", i+1)
		mockLog.AssertNumberOfCalls(t, "Close", 0)

		mockSM.AssertNumberOfCalls(t, "OpenConn", i+1)
	}
}
