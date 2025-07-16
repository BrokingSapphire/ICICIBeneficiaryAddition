package internal

import (
	"github.com/brokingSapphire/SapphireICICI/internal/env"
	"github.com/brokingSapphire/SapphireICICI/internal/logger"
	"github.com/brokingSapphire/SapphireICICI/internal/middleware"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// setupAPIRoutes configures API routes
func setupAPIRoutes(router *gin.RouterGroup) {
	// Test routes
	router.GET("/test", func(c *gin.Context) {
		logger.Info("Test endpoint called")
		c.JSON(200, gin.H{
			"message":     "ICICI API Service is running",
			"environment": env.Env.Env,
			"api_path":    env.Env.APIPath,
		})
	})
}

// SetupApp configures and returns the Gin application
func SetupApp() *gin.Engine {
	if env.Env.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Create Gin router
	app := gin.New()

	// Security middleware
	app.Use(Security())

	// CORS middleware
	app.Use(CORS())

	// Logging middleware
	app.Use(middleware.RequestLogger())

	// Health check route
	app.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "OK",
			"service": "ICICI API Service",
		})
	})

	// API routes group
	api := app.Group(env.Env.APIPath)
	{
		setupAPIRoutes(api)
	}

	logger.InfoWithFields(logrus.Fields{
		"api_path": env.Env.APIPath,
	}, "API routes registered")

	// Error handling middleware
	app.Use(middleware.ErrorHandler())

	// Handle 404 routes
	app.NoRoute(middleware.NotFoundHandler())

	return app
}
