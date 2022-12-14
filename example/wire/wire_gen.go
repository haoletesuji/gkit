// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"example/pkg"
	"gkit"
)

// Injectors from wire.go:

func InitializeServer(name string) (*gkit.Server, error) {
	httpLogger := gkit.NewHttpLogger()
	engine, err := pkg.NewGinServer(httpLogger)
	if err != nil {
		return nil, err
	}
	httpServer := pkg.NewHttpServer(name, httpLogger, engine)
	runner := pkg.NewRunner(httpServer)
	closer := pkg.NewCloser()
	server := gkit.NewServer(name, runner, closer)
	return server, nil
}
