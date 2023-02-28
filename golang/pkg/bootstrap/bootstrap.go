package bootstrap

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func WaitForTermination(errsChan chan error) {
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-sigterm:
		fmt.Println("Received interruption signal, application will shutdown.")
	case err := <-errsChan:
		fmt.Println("Received a critical error, application will shutdown.", err)
	}
}
