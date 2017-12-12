package utils

import (
	"runtime"

	log "github.com/inconshreveable/log15"
)

func HandleError(err error) {
	if err == nil {
		return
	}
	pc, fn, line, _ := runtime.Caller(1)
	log.Error(err.Error(), "Process", pc, "Function", fn, "Line", line)
}
