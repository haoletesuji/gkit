//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package wire

import (
	"httpexp/pkg"
	gkit "gkit"

	"github.com/google/wire"
)

func InitializeServer(name string) (*gkit.Server, error) {
	wire.Build(
		gkit.NewHttpLogger,
		pkg.NewGinServer,
		pkg.NewRunner,
		pkg.NewCloser,
		pkg.NewHttpServer,
		gkit.NewServer,
	)
	return &gkit.Server{}, nil
}
