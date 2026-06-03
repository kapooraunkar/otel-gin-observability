package telemetry

import (
	nirikshaai "github.com/san-data-systems/niriksha-sdk-go"
	"go.opentelemetry.io/otel/metric"
)

// ProductsCreated tracks how many products have been created.
var ProductsCreated metric.Int64Counter

// InitMetrics registers custom application metrics.
func InitMetrics() error {

	// Get a meter used to create and record metrics.
	meter := nirikshaai.Meter("product-service")

	// Create a counter that increments whenever a new product is created.
	counter, err := meter.Int64Counter(
		"products.created",
		metric.WithDescription(
			"Total number of products created",
		),
	)

	if err != nil {
		return err
	}

	// Store the counter so it can be used throughout the application.
	ProductsCreated = counter

	return nil
}
