package bootstrap

import (
	"net"
	"net/http"
	"sync"
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

func CreateServer(handler http.Handler, errCh chan error) *http.Server {
	server := &http.Server{
		Handler: handler,
	}

	go func() {
		errCh <- server.Serve(getListener())
	}()

	return server
}
