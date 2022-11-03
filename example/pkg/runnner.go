package pkg

import (
	"context"

	gkit "gkit"
)

type Runner struct {
	httpServer gkit.HttpServer
}

func NewRunner(httpServer gkit.HttpServer) gkit.Runner {
	return &Runner{
		httpServer,
	}
}

func (r *Runner) Run() {
	r.httpServer.RegisterRoutes()
	r.httpServer.Run()
}

func (r *Runner) GracefulShutdown(ctx context.Context) error {
	return r.httpServer.GracefulShutdown(ctx)
}
