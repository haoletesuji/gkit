package gkit

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var (
	JWTAuthHeader = "Authorization"
	JaegerHeader  = "Trace-Id"
)

func MaxAllowed(n int64) gin.HandlerFunc {
	sem := make(chan struct{}, n)
	acquire := func() { sem <- struct{}{} }
	release := func() { <-sem }

	return func(c *gin.Context) {
		acquire()       // before request
		defer release() // after request
		c.Next()
	}
}

func LoggingMiddleware(logger HttpLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()

		// Process Request
		c.Next()

		// Stop timer
		duration := getDurationInMillseconds(start)

		entry := logger.WithFields(log.Fields{
			"duration_ms": duration,
			"method":      c.Request.Method,
			"path":        c.Request.RequestURI,
			"status":      c.Writer.Status(),
			"referrer":    c.Request.Referer(),
			"trace_id":    getTraceID(c),
		})

		if c.Writer.Status() >= 500 {
			entry.Error(c.Errors.String())
		} else {
			entry.Info("")
		}
	}
}

func ExtractTokenFromHeader(r *http.Request) string {
	bearToken := r.Header.Get(JWTAuthHeader)
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func CorsMiddleware() gin.HandlerFunc {
	config := cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", JWTAuthHeader},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}
	return cors.New(config)
}

func getTraceID(c *gin.Context) string {
	identifier := c.Request.Header.Get(JaegerHeader)
	vals := strings.Split(identifier, ":")
	if len(vals) == 4 {
		return vals[0]
	}
	return ""
}

func getDurationInMillseconds(start time.Time) float64 {
	end := time.Now()
	duration := end.Sub(start)
	milliseconds := float64(duration) / float64(time.Millisecond)
	rounded := float64(int(milliseconds*100+.5)) / 100
	return rounded
}
