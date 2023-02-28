package main

import (
	"github.com/cioti/monorepo/cms.api/internal/transport"
	"github.com/cioti/monorepo/pkg/approot"
	"github.com/cioti/monorepo/pkg/bootstrap"
)

var errCh = make(chan error)

func main() {
	logger := approot.Infra().Logger()

	logger.Info("logging test")
	handler := transport.CreateHandler()
	bootstrap.CreateServer(handler, errCh)

	bootstrap.WaitForTermination(errCh)
}
