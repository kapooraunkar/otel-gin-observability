package services

import (
	"context"
	"fmt"
	"time"

	"go.opentelemetry.io/otel"
)

// Product represents a product in the catalog.
type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// In-memory product store used for demonstration purposes.
// In a real application this would typically be a database.
var products = []Product{
	{
		ID:    "p1",
		Name:  "Widget A",
		Price: 9.99,
	},
	{
		ID:    "p2",
		Name:  "Widget B",
		Price: 19.99,
	},
}

// Tracer used for service-layer operations.
var tracer = otel.Tracer("product-service")

// GetAllProducts retrieves all products.
//
// Supports:
//   - latency simulation (2 second delay)
//   - failure simulation (fail=true)
//   - service-level tracing
func GetAllProducts(
	ctx context.Context,
	fail bool,
) ([]Product, error) {

	// Create a child span to measure time spent in the service layer.
	_, span := tracer.Start(
		ctx,
		"service.getProducts",
	)
	defer span.End()

	// Simulate a slow database or external service call.
	time.Sleep(
		2 * time.Second,
	)

	// Simulate a database failure for observability testing.
	if fail {
		return nil,
			fmt.Errorf(
				"database connection failed",
			)
	}

	return products, nil
}

// CreateProduct adds a new product to the in-memory store.
func CreateProduct(p Product) {
	products = append(products, p)
}
