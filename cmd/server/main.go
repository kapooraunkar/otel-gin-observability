package main

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"

	"otel-gin-observability/internal/config"
	"otel-gin-observability/internal/handlers"
	"otel-gin-observability/internal/telemetry"

	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func main() {

	// Load application configuration from environment variables.
	cfg := config.Load()

	// Initialize Niriksha/OpenTelemetry tracing, metrics and logs.
	// Returns a shutdown function used to flush telemetry before exit.
	shutdown, err := telemetry.Init(cfg)
	if err != nil {
		panic(err)
	}

	// Register custom application metrics.
	if err := telemetry.InitMetrics(); err != nil {
		panic(err)
	}

	// Ensure telemetry data is exported before the application stops.
	defer shutdown(context.Background())

	// Print active configuration for debugging and verification.
	fmt.Println("APP ENV:", cfg.AppEnv)
	fmt.Println("TELEMETRY MODE:", cfg.TelemetryMode)

	// Create a Gin router with built-in logger and recovery middleware.
	router := gin.Default()

	// Automatically create traces for incoming HTTP requests.
	router.Use(
		otelgin.Middleware(cfg.ServiceName),
	)

	// Health check endpoint used to verify service availability.
	router.GET(
		"/health",
		handlers.Health,
	)

	// Returns all products.
	router.GET(
		"/products",
		handlers.GetProducts,
	)

	// Creates a new product.
	router.POST(
		"/products",
		handlers.CreateProduct,
	)

	// Start the HTTP server on port 8000.
	if err := router.Run(":8000"); err != nil {
		panic(err)
	}
}
