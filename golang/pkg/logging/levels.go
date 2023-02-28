package logging

import "github.com/go-kit/kit/log/level"

type LogLevel string

const (
	// Debug log level.
	Debug LogLevel = "DEBUG"
	// Info log level.
	Info LogLevel = "INFO"
	// Warn log level.
	Warn LogLevel = "WARN"
	// Error log level.
	Error LogLevel = "ERROR"
	// Panic log level.
	Panic LogLevel = "PANIC"
)

var (
	levelMap = map[LogLevel]level.Option{
		Debug: level.AllowDebug(),
		Info:  level.AllowInfo(),
		Warn:  level.AllowWarn(),
		Error: level.AllowError(),
		Panic: level.AllowError(),
	}
	defaultLevelOption = level.AllowAll()
)

func toLevelOption(level LogLevel) level.Option {
	opt, ok := levelMap[level]
	if !ok {
		return defaultLevelOption
	}

	return opt
}
