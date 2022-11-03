package pkg

import (
	gkit "gkit"

	"github.com/gin-gonic/gin"
)

func NewGinServer(logger gkit.HttpLogger) (*gin.Engine, error) {
	gkit.InitLogger("debug")

	engine := gin.New()
	engine.Use(gkit.LoggingMiddleware(logger))
	engine.Use(gkit.CorsMiddleware())

	return engine, nil
}
