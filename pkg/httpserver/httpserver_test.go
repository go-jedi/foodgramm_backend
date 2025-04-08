package httpserver

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"os/signal"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-jedi/foodgrammm-backend/config"
)

func newTestHTTPServer() *HTTPServer {
	gin.SetMode(gin.TestMode)

	cfg := config.HTTPServerConfig{
		Host:    "127.0.0.1",
		Port:    8080,
		GinMode: gin.TestMode,
		Cors: config.CorsConfig{
			AllowOrigins:        []string{"*"},
			AllowMethods:        []string{"GET", "POST"},
			AllowHeaders:        []string{"Content-Type"},
			ExposeHeaders:       []string{},
			MaxAge:              60,
			AllowCredentials:    true,
			AllowPrivateNetwork: false,
		},
	}

	hs, err := New(cfg)
	if err != nil {
		log.Fatalf("failed to create HTTP server: %v", err)
	}

	return hs
}

func TestPing(t *testing.T) {
	hs := newTestHTTPServer()

	req, err := http.NewRequest(http.MethodGet, "/ping", nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()
	hs.Engine.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, rr.Code)
	}

	var response map[string]string
	if err = json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Errorf("failed to unmarshal response body: %v", err)

	}

	if response["message"] != "pong" {
		t.Errorf("expected message 'pong', got '%s'", response["message"])
	}
}

func TestGracefulShutdown(t *testing.T) {
	hs := newTestHTTPServer()

	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", hs.port),
		Handler: hs.Engine,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			t.Logf("listen: %s\n", err)
		}
	}()

	quit <- os.Interrupt

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		t.Errorf("expected no error during shutdown, got: %v", err)
	}

	t.Log("server shutdown successfully")
}
