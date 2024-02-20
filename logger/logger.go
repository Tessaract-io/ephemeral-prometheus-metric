package logger

import (
	"log"
	"os"
)

var logger *log.Logger

func init() {
	if logger == nil {
		logger = log.New(os.Stdout, "ephemeral-prometheus-metric: ", log.LstdFlags)
	}
}

func GetLogger() *log.Logger {
	return logger
}
