package pkg

import (
	"context"
	"fmt"
	gkit "gkit"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	name      string
	logger    gkit.HttpLogger
	ginEngine *gin.Engine
	httpPort  string
	server    *http.Server
}

func NewHttpServer(name string, logger gkit.HttpLogger, ginEngine *gin.Engine) gkit.HttpServer {
	return &HttpServer{
		name:      name,
		logger:    logger,
		ginEngine: ginEngine,
		httpPort:  "3000",
	}
}

func (s *HttpServer) RegisterRoutes() {
	chatGroup := s.ginEngine.Group("/api")
	{
		chatGroup.GET("/health", s.HealthCheck)
		chatGroup.GET("/error", s.Error)
		chatGroup.GET("/success", s.Success)
		chatGroup.GET("/success_paging", s.SuccessPaging)

	}
}

func (s *HttpServer) Run() {
	go func() {
		addr := fmt.Sprintf(":%s", s.httpPort)
		s.server = &http.Server{
			Addr:    addr,
			Handler: s.ginEngine,
		}

		s.logger.Infoln(s.name, "server listening on", addr)
		err := s.server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			s.logger.Fatal(err)
		}
	}()

}

func (s *HttpServer) GracefulShutdown(ctx context.Context) error {
	err := s.server.Shutdown(ctx)
	if err != nil {
		return err
	}

	return nil
}
