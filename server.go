package gkit

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"
)

type Server struct {
	name     string
	runner   Runner
	closer   Closer
	migrator Migrator
}

func NewServer(name string, runner Runner, closer Closer, migrator Migrator) *Server {
	return &Server{
		name, runner, closer, migrator,
	}
}

func (s *Server) Serve() {
	err := s.migrator.Migrate()
	if err != nil {
		log.Fatal(err)
	}

	s.runner.Run()

	done := make(chan bool, 1)
	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
		<-sig

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		s.GracefulShutdown(ctx, done)
	}()
	<-done
}

func (s *Server) GracefulShutdown(ctx context.Context, done chan bool) {
	err := s.runner.GracefulShutdown(ctx)
	if err != nil {
		log.Error(err)
	}

	if err = s.closer.Close(); err != nil {
		log.Error(err)
	}

	log.Info("gracefully shutdowned")
	done <- true
}
