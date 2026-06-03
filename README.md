# Gin Observability Demo with Niriksha SDK

A Go application built with Gin, OpenTelemetry, and the Niriksha SDK to demonstrate observability concepts including tracing, metrics, logs, latency analysis, error tracking, and traffic generation.

## Features

* Gin HTTP API
* Automatic tracing via `otelgin`
* Custom spans in handlers and services
* Span events
* Application logs
* Custom metrics
* Error simulation
* Latency simulation
* Traffic generation script
* Docker support
* Niriksha SDK integration

---

## Prerequisites

* Go 1.26.3
* Python 3.x
* Docker (optional)
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

## Run Locally

Install dependencies:

```bash
go mod tidy
```

Start the application:

```bash
go run cmd/server/main.go
```

Application URL:

```text
http://localhost:8000
```

---

## API Endpoints

### Health Check

```http
GET /health
```

### List Products

```http
GET /products
```

### Create Product

```http
POST /products
```

### Simulate Failure

```http
GET /products?fail=true
```

---

## Observability Features

### Custom Spans

Handler span:

```text
products.list
products.create
```

Service span:

```text
service.getProducts
```

### Span Events

```text
products fetched successfully
product created
database failure
```

### Logs

Examples:

```text
products requested, count=2
product created: id=p123 name=Widget price=19.99
failed to fetch products: database connection failed
```

### Metrics

Counter:

```text
products.created
```

Counter:

```text
products.requested
```

---

## Traffic Generator

Install dependency:

```bash
pip install requests
```

Run:

```bash
python scripts/generate_traffic.py
```

The script continuously generates:

```text
GET /products
POST /products
GET /products?fail=true
```

to produce traces, logs, metrics, and errors.

---

## Docker

Build image:

```bash
docker build -t otel-gin-observability .
```

Run container:

```bash
docker run --env-file .env -p 8000:8000 otel-gin-observability
```

---

## Docker Validation Notes

Testing performed using:

```text
Go 1.26.3
Docker Desktop
Niriksha SDK v0.0.8
```

Observed behavior:

* Application starts successfully inside Docker.
* `/health` responds correctly.
* `/products` responds correctly.

---

## Project Goals

This project demonstrates:

* Distributed tracing
* Custom spans
* Span events
* Metrics
* Logs
* Error tracking
* Latency analysis
* Continuous telemetry generation
* OpenTelemetry instrumentation
* Docker deployment validation
