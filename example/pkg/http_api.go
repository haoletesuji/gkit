package pkg

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *HttpServer) HealthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
