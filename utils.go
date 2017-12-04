package utils

import (
	"os"
	"os/signal"
	"syscall"

	log "github.com/inconshreveable/log15"
)

// TillShutdown : This will ensure that the API backend is
// running without issues and does not terminate on errors.
// Purely depends on  what kind of SIGTERM is send to the
// application to terminate the threads.
func TillShutdown() {
	cs := make(chan os.Signal, 1)
	signal.Notify(cs, os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	for {
		select {
		case sig := <-cs:
			log.Debug("Captured Signal", "signal", sig)
			log.Error("Shutting Down...")
			return
		}
	}
}
