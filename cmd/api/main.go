package main

// @title User Management API
// @version 1.0
// @description REST API documentation for the Go Gin user management service.
// @BasePath /api/v1
// @schemes http https
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name X-API-Key
// @description API key required by the server middleware.
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Bearer token in the format: Bearer {token}.

import (
	"path/filepath"
	"user-management-api/internal/app"
	"user-management-api/internal/config"
	"user-management-api/internal/utils"
	"user-management-api/pkg/logger"

	"github.com/joho/godotenv"
)

func main() {
	rootDir := utils.MustGetWorkingDir()

	logFile := filepath.Join(rootDir, "internal/logs/app.log")

	logger.InitLogger(logger.LoggerConfig{
		Level:      "info",
		Filename:   logFile,
		MaxSize:    1,
		MaxBackups: 5,
		MaxAge:     5,
		Compress:   true,
		IsDev:      utils.GetEnv("APP_EVN", "development"),
	})

	if err := godotenv.Load(filepath.Join(rootDir, ".env")); err != nil {
		logger.Log.Warn().Msg("⚠️ No .env file found")
	} else {
		logger.Log.Info().Msg("✅ Loaded successfully .env in api proccess")
	}

	// Initialize configuration
	cfg := config.NewConfig()

	// Initialize application
	application, err := app.NewApplication(cfg)
	if err != nil {
		logger.Log.Fatal().Err(err).Msg("Failed to initialize application")
	}

	// Start server
	if err := application.Run(); err != nil {
		logger.Log.Fatal().Err(err).Msg("Application run failed")
	}
}
