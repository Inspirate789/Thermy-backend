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
	Name string
	Type MessageType
	Msg  string
}

type Logger interface {
	Open(serviceName string) error
	Print(r LogRecord)
	Close()
}
