package handlers

import (
	"log"
	"otel-gin-observability/internal/services"
	"otel-gin-observability/internal/telemetry"

	"github.com/gin-gonic/gin"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

// GetProducts returns all available products.
//
// Supports:
//   - latency simulation
//   - failure simulation (?fail=true)
//   - tracing via custom spans
func GetProducts(c *gin.Context) {

	// Create a span for the product listing operation.
	_, span := tracer.Start(
		c.Request.Context(),
		"products.list",
	)
	defer span.End()

	// Simulate a service failure using:
	// GET /products?fail=true
	fail := c.Query("fail") == "true"

	// Fetch products from the service layer.
	products, err := services.GetAllProducts(
		c.Request.Context(),
		fail,
	)

	telemetry.ProductsRequested.Add(
		c.Request.Context(),
		1,
	)

	if err != nil {

		span.AddEvent(
			"database failure",
		)

		log.Printf(
			"failed to fetch products: %v",
			err,
		)

		// Record the error in the active trace.
		span.RecordError(err)

		c.JSON(
			500,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	// Add useful business metadata to the trace.
	span.SetAttributes(
		attribute.Int("products.count", len(products)),
	)

	log.Printf(
		"products requested, count=%d",
		len(products),
	)

	span.AddEvent(
		"products fetched successfully",
	)

	c.JSON(
		200,
		gin.H{
			"products": products,
		},
	)
}

// CreateProduct creates a new product and records telemetry data.
func CreateProduct(c *gin.Context) {

	var p services.Product

	// Create a span for the product creation workflow.
	_, span := tracer.Start(
		c.Request.Context(),
		"products.create",
	)

	defer span.End()

	// Parse and validate the incoming JSON request body.
	if err := c.ShouldBindJSON(&p); err != nil {

		// Record validation errors in the trace.
		span.RecordError(err)

		c.JSON(
			400,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	// Attach product metadata to the trace.
	span.SetAttributes(
		attribute.String("product.id", p.ID),
		attribute.String("product.name", p.Name),
	)

	// Store the product using the service layer.
	services.CreateProduct(p)

	span.AddEvent(
		"product created",
	)

	log.Printf(
		"product created: id=%s name=%s price=%.2f",
		p.ID,
		p.Name,
		p.Price,
	)

	// Increment the custom metric whenever a product is created.
	telemetry.ProductsCreated.Add(
		c.Request.Context(),
		1,
	)

	c.JSON(
		201,
		p,
	)
}

// Tracer used for product-related operations.
var tracer = otel.Tracer("product-handler")
