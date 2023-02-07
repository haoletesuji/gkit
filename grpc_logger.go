package gkit

import (
	log "github.com/sirupsen/logrus"
)

type GrpcLogger struct {
	*log.Entry
}

func NewGrpcLogger() GrpcLogger {
	return GrpcLogger{log.WithField("type", "grpc")}
}
