package logger

import "github.com/stretchr/testify/mock"

type MockLogger struct {
	mock.Mock
}

func (m *MockLogger) Open(serviceName string) error {
	args := m.Called(serviceName)
	return args.Error(0)
}

func (m *MockLogger) Print(r LogRecord) {
	m.Called(r)
}

func (m *MockLogger) Close() {
	m.Called()
}
