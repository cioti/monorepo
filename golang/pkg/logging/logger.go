package logging

import (
	"github.com/go-kit/log"
)

const Message = "message"

// Logger implements go-kit Logger and adds wrappers to log with levels.
type Logger interface {
	log.Logger
	// Info writes logs with info level.
	Info(keyValues ...interface{})
	// Debug writes logs with debug level.
	Debug(keyValues ...interface{})
	// Warn writes logs with warn level.
	Warn(keyValues ...interface{})
	// Error writes logs with error level.
	Error(keyValues ...interface{})
}
