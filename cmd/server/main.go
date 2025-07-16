package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/brokingSapphire/SapphireICICI/internal"
	"github.com/brokingSapphire/SapphireICICI/internal/env"
	"github.com/brokingSapphire/SapphireICICI/internal/logger"
	"github.com/sirupsen/logrus"
)

func main() {
	// Load environment configuration
	env.Load()

	// Initialize logger
	logger.Init()

	// Setup the Gin application using your new app.go
	router := internal.SetupApp()

	// Create HTTP server with proper configuration
	server := &http.Server{
		Addr:    ":" + env.Env.Port,
		Handler: router,
		// Configure timeouts for production
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// Start server in a goroutine so it doesn't block
	go func() {
		logger.InfoWithFields(logrus.Fields{
			"port":        env.Env.Port,
			"environment": env.Env.Env,
			"api_path":    env.Env.APIPath,
		}, "Server starting")

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("Failed to start server: ", err)
		}
	}()

	// Setup graceful shutdown
	setupGracefulShutdown(server)
}

// setupGracefulShutdown handles graceful server shutdown
func setupGracefulShutdown(server *http.Server) {
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Block until a signal is received
	sig := <-quit
	logger.InfoWithFields(logrus.Fields{
		"signal": sig.String(),
		"time":   time.Now().Format(time.RFC3339),
	}, "Received shutdown signal, starting graceful shutdown")

	// Create a context with timeout for shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Attempt graceful shutdown
	if err := server.Shutdown(ctx); err != nil {
		logger.Error("Server forced to shutdown: ", err)
		os.Exit(1)
	}

	logger.Info("Server shutdown complete")
}
