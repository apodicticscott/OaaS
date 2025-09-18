# Start from official Go image
FROM golang:1.25 AS builder

WORKDIR /app

# Install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Build Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o ontology-service ./cmd/server

# --- Runtime container ---
FROM alpine:3.18

WORKDIR /root/

# Copy binary
COPY --from=builder /app/ontology-service .

EXPOSE 8080

CMD ["./ontology-service"]
