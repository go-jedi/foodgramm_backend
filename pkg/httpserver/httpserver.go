package httpserver

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-jedi/foodgrammm-backend/config"
	_ "github.com/go-jedi/foodgrammm-backend/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const (
	defaultHost     = "127.0.0.1"
	defaultPort     = 50050
	readTimeoutSec  = 10
	writeTimeoutSec = 10
	idleTimeout     = 120
	maxHeaderBytes  = 1 << 20
)

type HTTPServer struct {
	Engine *gin.Engine

	host    string
	port    int
	ginMode string
}

func (hs *HTTPServer) init() error {
	if hs.host == "" {
		hs.host = defaultHost
	}

	if hs.port == 0 {
		hs.port = defaultPort
	}

	if hs.ginMode == "" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(hs.ginMode)
	}

	engine := gin.New()

	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	hs.Engine = engine

	return nil
}

func NewHTTPServer(cfg config.HTTPServerConfig) (*HTTPServer, error) {
	hs := &HTTPServer{
		host:    cfg.Host,
		port:    cfg.Port,
		ginMode: cfg.GinMode,
	}

	if err := hs.init(); err != nil {
		return nil, err
	}

	hs.initSwagger()
	hs.ping()

	return hs, nil
}

// Start http server.
func (hs *HTTPServer) Start() error {
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", hs.port),
		Handler:        hs.Engine,                     // Обработчик запросов — это экземпляр Gin
		ReadTimeout:    readTimeoutSec * time.Second,  // Максимальное время ожидания запроса
		WriteTimeout:   writeTimeoutSec * time.Second, // Максимальное время для обработки ответа
		IdleTimeout:    idleTimeout * time.Second,     // Время бездействия перед закрытием соединения
		MaxHeaderBytes: maxHeaderBytes,                // Максимальный размер заголовков
	}

	if err := hs.gracefulStop(s); err != nil {
		return fmt.Errorf("httpserver graceful stop error: %w", err)
	}

	return nil
}

func (hs *HTTPServer) initSwagger() {
	hs.Engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

// @Summary Ping
// @Description Check server status
// @Tags Health
// @Produce json
// @Success 200 {object} map[string]string "message: pong"
// @Router /ping [get]
func (hs *HTTPServer) ping() {
	hs.Engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}

// gracefulStop server with graceful shutdown.
func (hs *HTTPServer) gracefulStop(s *http.Server) error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-quit
	log.Println("shutting down server...")

	const ctxSec = 5
	ctx, cancel := context.WithTimeout(context.Background(), ctxSec*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Printf("server forced to shutdown: %v", err)
		return err
	}

	log.Println("server exiting")

	return nil
}
