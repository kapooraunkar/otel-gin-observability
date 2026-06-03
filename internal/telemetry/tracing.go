package telemetry

import (
	"context"
	"fmt"

	"otel-gin-observability/internal/config"

	nirikshaai "github.com/san-data-systems/niriksha-sdk-go"
)

// Init configures and initializes Niriksha/OpenTelemetry.
//
// Returns:
//   - a shutdown function used to flush telemetry before application exit
//   - an error if initialization fails
func Init(cfg config.Config) (func(context.Context) error, error) {

	var endpoint string

	switch cfg.TelemetryMode {

	case "grpc-direct":
		endpoint = cfg.OTLPGrpcEndpoint

	case "http-direct":
		endpoint = cfg.OTLPHttpEndpoint

	default:
		return nil,
			fmt.Errorf(
				"unsupported telemetry mode: %s",
				cfg.TelemetryMode,
			)
	}

	// Configure the Niriksha SDK with application metadata
	// and telemetry export settings.
	return nirikshaai.Init(
		context.Background(),
		nirikshaai.Options{

			// Niriksha dashboard URL.
			Endpoint: "https://app.niriksha.ai",

			// OTLP ingestion endpoint used for exporting traces,
			// metrics and logs.
			OTLPEndpoint: endpoint,

			// API key used to authenticate with Niriksha.
			APIKey: cfg.APIKey,

			// Service name displayed in telemetry data.
			ServiceName: cfg.ServiceName,

			// Environment label (development, staging, production).
			Environment: cfg.AppEnv,

			// Enable metric collection and export.
			EnableMetrics: true,

			// Enable log collection and export.
			EnableLogs: true,
		},
	)
}
