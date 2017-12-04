package utils

import (
	"net/http"
	"net/http/httputil"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/ankitforcode/go-utils/config"
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

func LoggerMiddleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, r)
			if config.Config.Server.Loglevel == "debug" {
				requestDump, _ := httputil.DumpRequest(r, true)
				log.Debug(r.URL.String(), "Request", string(requestDump))
			} else {
				log.Info(r.URL.String(), "Method", r.Method, "Len", r.ContentLength, "Proto", r.Proto,
					"Latency", time.Since(start), "Raddr", getIpAddress(r))
			}
		})
	}
}

func LoggerInit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Info(r.URL.String(), "Method", r.Method, "Len", r.ContentLength, "Proto", r.Proto,
			"Latency", time.Since(start), "Raddr", getIpAddress(r))
	})
}

func getIpAddress(r *http.Request) string {
	hdr := r.Header
	hdrRealIP := hdr.Get("X-Real-Ip")
	hdrForwardedFor := hdr.Get("X-Forwarded-For")
	if hdrRealIP == "" && hdrForwardedFor == "" {
		return ipAddrFromRemoteAddr(r.RemoteAddr)
	}
	if hdrForwardedFor != "" {
		// X-Forwarded-For is potentially a list of addresses separated with ","
		parts := strings.Split(hdrForwardedFor, ",")
		for i, p := range parts {
			parts[i] = strings.TrimSpace(p)
		}
		// TODO: should return first non-local address
		return parts[0]
	}
	return hdrRealIP
}

func ipAddrFromRemoteAddr(s string) string {
	idx := strings.LastIndex(s, ":")
	if idx == -1 {
		return s
	}
	return s[:idx]
}

// FaviconHandler :
func FaviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/favicon.ico")
}
