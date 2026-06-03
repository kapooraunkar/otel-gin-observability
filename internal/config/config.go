package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Config stores all application configuration loaded from environment variables.
type Config struct {
	APIKey           string
	AppEnv           string
	TelemetryMode    string
	ServiceName      string
	OTLPGrpcEndpoint string
	OTLPHttpEndpoint string
}

// Load reads environment variables and returns a populated Config struct.
func Load() Config {

	// Load variables from a local .env file into the environment.
	// If no .env file exists, values can still come from system environment variables.
	err := godotenv.Load()
	if err != nil {
		println("No .env file found")
	}

	return Config{

		// Niriksha API key used for authentication.
		APIKey: os.Getenv("NIRIKSHA_API_KEY"),

		// Application environment (development, staging, production).
		AppEnv: os.Getenv("APP_ENV"),

		// Telemetry export mode (grpc-direct, http-direct, etc.).
		TelemetryMode: os.Getenv("TELEMETRY_MODE"),

		// Service name shown in traces, metrics and logs.
		ServiceName: os.Getenv("SERVICE_NAME"),

		// OTLP gRPC ingestion endpoint.
		OTLPGrpcEndpoint: os.Getenv("OTLP_GRPC_ENDPOINT"),

		// OTLP HTTP ingestion endpoint.
		OTLPHttpEndpoint: os.Getenv("OTLP_HTTP_ENDPOINT"),
	}
}
