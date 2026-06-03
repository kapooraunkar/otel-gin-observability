# Gin Observability Demo with Niriksha SDK

A Go application built with Gin and the Niriksha SDK to demonstrate observability concepts including distributed tracing, metrics, logs, latency analysis, error tracking, and continuous traffic generation using OpenTelemetry.

---

## Features

* Gin-based HTTP service
* Automatic request tracing using `otelgin`
* Custom handler and service spans
* Trace events
* Structured application logs
* Custom metrics
* Error simulation
* Latency simulation
* Continuous traffic generation script
* Environment-based configuration
* Niriksha SDK integration

---

## Project Structure

```text
cmd/
└── server/
    └── main.go

internal/
├── config/
│   └── config.go
├── handlers/
│   ├── health.go
│   └── product.go
├── services/
│   └── product_service.go
└── telemetry/
    ├── tracing.go
    └── metrics.go

scripts/
└── generate_traffic.py
```

---

## Prerequisites

* Go 1.22+
* Python 3.x
* Niriksha API Key

---

## Environment Variables

Create a `.env` file:

```env
NIRIKSHA_API_KEY=your_api_key

APP_ENV=development

SERVICE_NAME=otel-gin-observability

TELEMETRY_MODE=grpc-direct

OTLP_GRPC_ENDPOINT=grpc-ingest.niriksha.ai:443

OTLP_HTTP_ENDPOINT=grpc-ingest.niriksha.ai:443
```

---

## Installation

Install dependencies:

```bash
go mod tidy
```

---

## Running the Application

Start the server:

```bash
go run cmd/server/main.go
```

Application runs on:

```text
http://localhost:8000
```

---

## API Endpoints

### Health Check

```http
GET /health
```

Response:

```json
{
  "status": "ok"
}
```

---

### Get Products

```http
GET /products
```

Returns all available products.

---

### Create Product

```http
POST /products
```

Request:

```json
{
  "id": "p3",
  "name": "Widget C",
  "price": 29.99
}
```

Response:

```json
{
  "id": "p3",
  "name": "Widget C",
  "price": 29.99
}
```

---

## Observability Scenarios

### Latency Simulation

```http
GET /products
```

The service layer introduces an intentional 2-second delay.

Trace flow:

```text
HTTP Request
    ↓
products.list
    ↓
service.getProducts
```

This demonstrates how latency can be identified through traces.

---

### Error Simulation

```http
GET /products?fail=true
```

Response:

```json
{
  "error": "database connection failed"
}
```

This generates:

* Error traces
* Error events
* Error logs

---

### Trace Events

The application records custom trace events including:

```text
products fetched successfully

database failure

product created
```

These events appear within spans and provide additional execution context.

---

### Application Logs

Examples:

```text
products requested, count=10

product created: id=p1234 name=Product-15 price=42.50

failed to fetch products: database connection failed
```

Logs help correlate application behavior with traces and metrics.

---

## Metrics

### products.created

Counter metric that increments whenever a new product is created.

Triggered by:

```http
POST /products
```

---

### products.requested

Counter metric that increments whenever the products endpoint is requested.

Triggered by:

```http
GET /products
```

---

## Traffic Generator

A traffic generation script is included to continuously generate:

* Traces
* Metrics
* Logs
* Errors

Install Python dependency:

```bash
pip install requests
```

Run:

```bash
python scripts/generate_traffic.py
```

The script continuously performs:

```text
GET /products

POST /products

GET /products?fail=true (occasionally)
```

This allows observability dashboards to receive a constant stream of telemetry data.

---

## Telemetry Modes

Supported modes:

* grpc-direct
* http-direct

Current SDK examples use:

```text
grpc-ingest.niriksha.ai:443
```

for telemetry export.

---

## Tracing

Automatic request tracing is enabled through:

```go
otelgin.Middleware(...)
```

Custom spans include:

```text
products.list

products.create

service.getProducts
```

These spans help identify:

* Request latency
* Service bottlenecks
* Failures
* Business operations

---

## What This Demo Demonstrates

* Distributed tracing
* Custom spans
* Trace events
* Application logs
* Custom metrics
* Error tracking
* Latency analysis
* Continuous traffic generation
* Niriksha SDK integration
* OpenTelemetry instrumentation

```
```
