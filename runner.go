package gkit

import (
	"context"
)

type Runner interface {
	Run()
	GracefulShutdown(ctx context.Context) error
}
