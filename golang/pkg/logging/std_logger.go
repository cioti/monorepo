package logging

import (
	"io"
	"os"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

type stdLogger struct {
	stdOutLogger log.Logger
	stdErrLogger log.Logger
}

func NewLogger(name string, level LogLevel) Logger {
	levelOpt := toLevelOption(level)

	stdOutLogger := newFilteredLogger(name, os.Stdout, levelOpt)
	stdErrLogger := newFilteredLogger(name, os.Stderr, toLevelOption(Warn))

	return newStdLogger(stdOutLogger, stdErrLogger)
}

func (log *stdLogger) Log(values ...interface{}) error {
	return log.stdOutLogger.Log(values...)
}

func (log *stdLogger) Info(keyValues ...interface{}) {
	_ = level.Info(log.stdOutLogger).Log(keyValues...)
}

func (log *stdLogger) Debug(keyValues ...interface{}) {
	_ = level.Debug(log.stdOutLogger).Log(keyValues...)
}

func (log *stdLogger) Warn(keyValues ...interface{}) {
	_ = level.Warn(log.stdErrLogger).Log(keyValues...)
}

func (log *stdLogger) Error(keyValues ...interface{}) {
	_ = level.Error(log.stdErrLogger).Log(keyValues...)
}

func newStdLogger(stdOutLogger log.Logger, stdErrLogger log.Logger) *stdLogger {
	return &stdLogger{
		stdOutLogger: stdOutLogger,
		stdErrLogger: stdErrLogger,
	}
}

func newFilteredLogger(name string, output io.Writer, options ...level.Option) log.Logger {
	return level.NewFilter(
		log.With(
			log.NewJSONLogger(output),
			"timestamp", log.DefaultTimestampUTC,
			"loggerName", name,
		), options...,
	)
}
