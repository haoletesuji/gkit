package gkit

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestMaxAllowed(t *testing.T) {
	handler := MaxAllowed(2)

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	handler(ctx)
}

func TestExtractTokenFromHeader(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "http://example.com", nil)
	req.Header = map[string][]string{
		"Accept-Encoding": {"gzip, deflate"},
		"Accept-Language": {"en-us"},
		"Authorization":   {"Bearer TokenToken"},
	}

	token := ExtractTokenFromHeader(req)
	if token != "TokenToken" {
		t.Errorf("ExtractTokenFromHeader was incorrect, got: %s, want: %s.", token, "TokenToken")
	}
}
