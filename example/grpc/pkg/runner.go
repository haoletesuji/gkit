package pkg

import (
	"context"

	gkit "gkit"
)

type Runner struct {
	grpcServer gkit.GrpcServer
}

func NewRunner(grpcServer gkit.GrpcServer) gkit.Runner {
	return &Runner{
		grpcServer: grpcServer,
	}
}

func (r *Runner) Run() {
	r.grpcServer.Register()
	r.grpcServer.Run()
}

func (r *Runner) GracefulShutdown(ctx context.Context) error {
	r.grpcServer.GracefulShutdown()
	return nil
}
