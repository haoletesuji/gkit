//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package wire

import (
	gkit "gkit"
	"httpexp/pkg"

	"github.com/google/wire"
)

func InitializeServer(name string) (*gkit.Server, error) {
	wire.Build(
		gkit.NewHttpLogger,
		pkg.NewGinServer,
		pkg.NewRunner,
		pkg.NewCloser,
		pkg.NewHttpServer,
		pkg.NewMigrator,
		gkit.NewServer,
	)
	return &gkit.Server{}, nil
}
