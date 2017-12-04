package server

import (
	"fmt"
	"net/http"

	"github.com/ankitforcode/go-utils/config"
	log "github.com/inconshreveable/log15"
)

type servers struct {
	Host string
	Port uint
}

// Run : Start the Server
func Run(httpHandlers http.Handler) (s servers) {
	s.Host = config.Config.Server.Host
	s.Port = config.Config.Server.Port
	log.Info("Starting Server On :", "address", s.Host, "port", s.Port)
	startServer(s, httpHandlers)
	return
}

func startServer(s servers, handler http.Handler) {
	if err := http.ListenAndServe(httpAddress(s), handler); err != nil {
		log.Error("Error", "error", err)
	}
}

func httpAddress(s servers) string {
	return s.Host + ":" + fmt.Sprintf("%d", s.Port)
}
