package logger

type MessageType int

const (
	Debug MessageType = iota
	Warning
	Error
)

func (t MessageType) String() string {
	return [...]string{"DEBUG", "WARNING", "ERROR"}[t]
}

type LogRecord struct {
	name string
	t    MessageType
	s    string
}

type Logger interface {
	Open() error
	Print(LogRecord) error
	Close() error
}
