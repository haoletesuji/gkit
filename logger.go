package common

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func InitLogger(mode string) {
	gin.SetMode(mode)
	gin.DefaultWriter = io.Writer(os.Stderr)
	log.SetOutput(os.Stderr)
	// log.SetLevel(log.TraceLevel)
	log.SetLevel(log.DebugLevel)
	// log.SetLevel(log.InfoLevel)
	// log.SetLevel(log.WarnLevel)
	// log.SetLevel(log.ErrorLevel)
	// log.SetLevel(log.FatalLevel)
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true, DisableColors: true})
}
