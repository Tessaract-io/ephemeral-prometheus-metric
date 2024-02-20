package utils

import (
	_logger "ephemeral-prometheus-metric/logger"
	"log"
)

var logger *log.Logger

func init() {
	logger = _logger.GetLogger()
}
