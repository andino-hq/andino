package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/andino-hq/andino/internal/infrastructure/config"
	"github.com/andino-hq/andino/internal/infrastructure/interfaces/http/router"
	"github.com/andino-hq/andino/internal/infrastructure/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(fmt.Sprintf("Failed to load config: %v", err))
	}

	log, err := logger.NewLogger(cfg.Logging)
	if err != nil {
		panic(fmt.Sprintf("Failed to setup logger: %v", err))
	}
	defer log.Sync()

	log.Info("Starting Andino server",
		zap.String("version", "1.0.0"),
		zap.String("environment", "development"))

	if cfg.Logging.Level == "debug" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := router.Setup(cfg, log)

	srv := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port),
		Handler:      r,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
	}

	go func() {
		log.Info("Server starting",
			zap.String("addr", srv.Addr),
			zap.String("message", "Andino is ready! 🏔️"))

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Failed to start server", zap.Error(err))
		}
	}()

	fmt.Printf(`
🏔️  ANDINO  🏔️
================================
✅ Server started at:   http://%s:%d
✅ Health check:        http://%s:%d/health
✅ Hello World:         http://%s:%d/hello
✅ API v1:              http://%s:%d/api/v1

Press Ctrl+C to stop the server...

`, cfg.Server.Host, cfg.Server.Port,
		cfg.Server.Host, cfg.Server.Port,
		cfg.Server.Host, cfg.Server.Port,
		cfg.Server.Host, cfg.Server.Port)

	// Wait for interrupt signal to gracefully shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown", zap.Error(err))
	}

	log.Info("Server exited gracefully")
}
