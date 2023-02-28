package approot

import (
	"sync"

	"github.com/cioti/monorepo/pkg/logging"
)

var (
	infraOnce sync.Once
	infra     InfraFactory
)

type InfraFactory interface {
	Logger() logging.Logger
}

type infraFactory struct {
	logger     logging.Logger
	loggerOnce sync.Once
}

func (f *infraFactory) Logger() logging.Logger {
	f.loggerOnce.Do(func() {
		f.logger = logging.NewLogger("logger", logging.Debug)
	})

	return f.logger
}

func Infra() InfraFactory {
	infraOnce.Do(func() {
		infra = &infraFactory{}
	})

	return infra
}
