# Build stage
FROM golang:1.26.3 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server ./cmd/server

# Runtime stage
FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 8000

CMD ["./server"]