package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"gitlab.mserv.wtf/navidrome-mix/pkg/config"
	"gitlab.mserv.wtf/navidrome-mix/pkg/db"
	"gitlab.mserv.wtf/navidrome-mix/pkg/navidrome"
	"gitlab.mserv.wtf/navidrome-mix/pkg/server"
)

func main() {
	// Parse command line flags
	configPath := flag.String("config", "./config/config.yaml", "Path to configuration file")
	flag.Parse()

	// Initialize logger
	logger := initLogger()
	defer logger.Sync()

	// Load configuration
	cfg, err := config.Load(*configPath)
	if err != nil {
		logger.Fatal("Failed to load configuration", zap.Error(err))
	}

	// Set log level from configuration
	setLogLevel(logger, cfg.Service.LogLevel)
	logger.Info("Starting Navidrome Mix Service", zap.String("version", "0.1.0"))

	// Initialize database connection
	dbClient, err := db.NewClient(cfg.Database.URI, cfg.Database.Username, cfg.Database.Password)
	if err != nil {
		logger.Fatal("Failed to connect to database", zap.Error(err))
	}
	logger.Info("Connected to database")

	// Initialize Navidrome client
	navClient := navidrome.NewClient(
		cfg.Navidrome.URL,
		cfg.Navidrome.APIPath,
		cfg.Navidrome.Username,
		cfg.Navidrome.Password,
	)
	logger.Info("Navidrome client initialized", zap.String("url", cfg.Navidrome.URL))

	// Create and initialize HTTP server
	srv := server.New(cfg, logger, dbClient, navClient)
	
	// Start HTTP server in a goroutine
	go func() {
		addr := fmt.Sprintf("%s:%d", cfg.Service.Host, cfg.Service.Port)
		logger.Info("Starting HTTP server", zap.String("address", addr))
		
		if err := srv.Run(addr); err != nil && err != http.ErrServerClosed {
			logger.Fatal("Server failed", zap.Error(err))
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	
	logger.Info("Shutting down server...")

	// Create a deadline for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("Server forced to shutdown", zap.Error(err))
	}

	// Close database connection
	if err := dbClient.Close(); err != nil {
		logger.Error("Error closing database connection", zap.Error(err))
	}

	logger.Info("Server exited gracefully")
}

func initLogger() *zap.Logger {
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	
	logger, err := config.Build()
	if err != nil {
		panic(fmt.Sprintf("failed to initialize logger: %v", err))
	}
	
	return logger
}

func setLogLevel(logger *zap.Logger, level string) {
	var atomicLevel zap.AtomicLevel
	
	switch level {
	case "debug":
		atomicLevel = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	case "info":
		atomicLevel = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	case "warn":
		atomicLevel = zap.NewAtomicLevelAt(zapcore.WarnLevel)
	case "error":
		atomicLevel = zap.NewAtomicLevelAt(zapcore.ErrorLevel)
	default:
		atomicLevel = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	}
	
	logger.Info("Log level set", zap.String("level", level))
}
