//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package wire

import (
	gkit "gkit"
	"grpcexp/pkg"

	"github.com/google/wire"
)

func InitializeServer(name string) (*gkit.Server, error) {
	wire.Build(
		gkit.NewGrpcLogger,
		pkg.NewCloser,
		pkg.NewRunner,
		pkg.NewGrpcServer,
		pkg.NewMigrator,
		gkit.NewServer,
	)
	return &gkit.Server{}, nil
}
