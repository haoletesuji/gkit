package common

import (
	log "github.com/sirupsen/logrus"
)

type HttpLogger struct {
	*log.Entry
}

func NewHttpLogger() HttpLogger {
	return HttpLogger{log.WithField("type", "http")}
}
