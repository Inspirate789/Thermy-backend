package authorization

import (
	"context"
	"github.com/Inspirate789/Thermy-backend/internal/domain/services/storage"
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"
)

func TestSession_Open(t *testing.T) {
	mockSM := new(storage.MockStorageManager)
	mockSM.On("OpenConn", mock.Anything, mock.Anything).Return(nil, "admin", nil)
	mockSM.On("CloseConn", mock.Anything).Return(nil)

	type args struct {
		sm      storage.StorageManager
		request *storage.AuthRequest
		ctx     context.Context
	}
	tests := []struct {
		name    string
		session *Session
		args    args
		want    uint64
		wantErr bool
	}{
		{
			name:    "Simple positive test",
			session: NewSession(),
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
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.session.Open(tt.args.sm, tt.args.request, tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Open() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Open() got = %v, want %v", got, tt.want)
			}

			mockSM.AssertNumberOfCalls(t, "OpenConn", i+1)
		})
	}
}

func TestSession_Close(t *testing.T) {
	mockSM := new(storage.MockStorageManager)
	mockSM.On("OpenConn", mock.Anything, mock.Anything).Return(nil, "admin", nil)
	mockSM.On("CloseConn", mock.Anything).Return(nil)

	type args struct {
		sm storage.StorageManager
	}
	tests := []struct {
		name    string
		session *Session
		args    args
		wantErr bool
	}{
		{
			name:    "Simple positive test",
			session: NewSession(),
			args: args{
				sm: mockSM,
			},
			wantErr: false,
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.session.Close(tt.args.sm); (err != nil) != tt.wantErr {
				t.Errorf("Close() error = %v, wantErr %v", err, tt.wantErr)
			}

			mockSM.AssertNumberOfCalls(t, "CloseConn", i+1)
		})
	}
}

func TestSession_GetAuthData(t *testing.T) {
	tests := []struct {
		name    string
		session *Session
		want    *storage.AuthRequest
	}{
		{
			name: "Simple positive test",
			session: &Session{
				authData: &storage.AuthRequest{
					Username: "initial_admin",
					Password: "12345",
				},
			},
			want: &storage.AuthRequest{
				Username: "initial_admin",
				Password: "12345",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.session.GetAuthData(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAuthData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSession_GetConn(t *testing.T) {
	tests := []struct {
		name    string
		session *Session
		want    storage.ConnDB
	}{
		{
			name:    "Simple positive test",
			session: &Session{connDB: "conn"},
			want:    "conn",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.session.GetConn(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetConn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSession_GetRole(t *testing.T) {
	tests := []struct {
		name    string
		session *Session
		want    string
	}{
		{
			name:    "Simple positive test",
			session: &Session{role: "admin"},
			want:    "admin",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.session.GetRole(); got != tt.want {
				t.Errorf("GetRole() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSession_GetToken(t *testing.T) {
	tests := []struct {
		name    string
		session *Session
		want    uint64
	}{
		{
			name:    "Simple positive test",
			session: &Session{token: 10063865700249539947},
			want:    10063865700249539947,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.session.GetToken(); got != tt.want {
				t.Errorf("GetToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
