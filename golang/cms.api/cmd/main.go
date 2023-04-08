package main

import (
	"github.com/cioti/monorepo/cms.api/service/app"
	"github.com/cioti/monorepo/cms.api/service/infra"
	"github.com/cioti/monorepo/cms.api/service/transport"
	"github.com/cioti/monorepo/cms.api/service/transport/http"
	"github.com/cioti/monorepo/pkg/approot"
	"github.com/cioti/monorepo/pkg/bootstrap"
)

var errCh = make(chan error)

func main() {
	logger := approot.Infra().Logger()

	repo := infra.NewProjectRepository(*approot.Infra().Mongo().Database("cms"))
	svc := app.NewProjectService(repo)
	endpoints := transport.NewEndpoints(svc)
	handler := http.CreateHandler(endpoints, logger)
	bootstrap.CreateServer(handler, logger, errCh)

	bootstrap.WaitForTermination(errCh)
}
