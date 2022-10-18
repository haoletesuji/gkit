package gkit

import (
	"context"
)

type HttpServer interface {
	RegisterRoutes()
	Run()
	GracefulShutdown(ctx context.Context) error
}
