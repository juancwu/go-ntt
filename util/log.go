package util

import (
	"os"

	log "github.com/charmbracelet/log"
)

var _log *log.Logger

func InitLog() {
	if _log != nil {
		return
	}

	_log = NewLogger()
}

func NewLogger() *log.Logger {
	l := log.InfoLevel
	envLevel := os.Getenv("GONTT_LOG_LEVEL")
	if envLevel == "debug" {
		l = log.DebugLevel
	}

	return log.NewWithOptions(os.Stdout, log.Options{Level: l, ReportTimestamp: true, ReportCaller: envLevel == "debug"})
}

func Log() *log.Logger {
	return _log
}
