package bootstrap

import (
	"fmt"
	"net"
	"net/http"
	"sync"

	"github.com/cioti/monorepo/pkg/logging"
)

var (
	serviceTCPListener net.Listener
	onceTCPListener    sync.Once
)

func getServiceAddress() string {
	// todo get from env variable port
	return ":8080"
}

func getListener() net.Listener {
	onceTCPListener.Do(func() {
		listener, err := net.Listen("tcp", getServiceAddress())
		if err != nil {
			panic(err)
		}
		serviceTCPListener = listener
	})

	return serviceTCPListener
}

func CreateServer(handler http.Handler, logger logging.Logger, errCh chan error) *http.Server {
	server := &http.Server{
		Handler: handler,
	}

	listener := getListener()
	go func() {
		logger.Info(logging.Message, fmt.Sprintf("Start HTTP listening on: %s", listener.Addr()))
		errCh <- server.Serve(listener)
	}()

	return server
}
