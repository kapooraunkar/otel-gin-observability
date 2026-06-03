# Gin Observability Demo with Niriksha SDK

A sample Go application built with Gin and the Niriksha SDK to demonstrate distributed tracing, metrics, error tracking, and latency analysis using OpenTelemetry.

## Features

* Gin HTTP server
* Automatic request tracing using `otelgin`
* Custom spans for handlers and services
* Custom metrics (`products.created`)
* Error tracking
* Latency simulation
* Service layer instrumentation
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
```

---

## Prerequisites

* Go 1.22+
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

## Run

```bash
go mod tidy
go run cmd/server/main.go
```

Application runs on:

```text
http://localhost:8000
```

---

## Endpoints

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

### List Products

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

The service layer intentionally introduces a 2-second delay to demonstrate latency analysis in traces.

Trace flow:

```text
products.list
    ↓
service.getProducts
```

---

### Failure Simulation

```http
GET /products?fail=true
```

Returns:

```json
{
  "error": "database connection failed"
}
```

This can be used to demonstrate error tracking and trace debugging.

---

## Metrics

### products.created

Counter metric that increments every time a product is successfully created.

Triggered by:

```http
POST /products
```

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

Custom spans are created for:

* products.list
* products.create
* service.getProducts

These spans help identify latency, failures, and service-level behavior.

---

## What This Demo Shows

* Request tracing
* Service tracing
* Error recording
* Custom metrics
* Latency analysis
* Telemetry integration using Niriksha SDK
