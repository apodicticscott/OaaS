APP_NAME=ontology-service
APP_PORT?=8080
DB_CONTAINER=ontology_db
DC=docker-compose -f docker/docker-compose.yml

# Default target
.DEFAULT_GOAL := help

## Run the Go server locally
run:
	@echo "🚀 Running locally on :$(APP_PORT)"
	go run cmd/server/main.go

## Run gqlgen to regenerate code from schema
gqlgen:
	@echo "⚙️  Running gqlgen..."
	go run github.com/99designs/gqlgen generate

## Build the Go binary
build:
	@echo "🔨 Building $(APP_NAME)"
	go build -o bin/$(APP_NAME) ./cmd/server

## Run docker-compose stack (Go + Postgres + pgAdmin)
docker-up:
	$(DC) up --build -d

## Stop docker-compose stack
docker-down:
	$(DC) down

## Show logs from Go service
logs:
	$(DC) logs -f ontology-service

## Run psql inside the DB container
psql:
	docker exec -it $(DB_CONTAINER) psql -U postgres -d ontology

## Run database migrations (requires migrate CLI or similar)
migrate-up:
	@echo "⚡ Running migrations..."
	#migrate -path db/migrations -database "postgres://postgres:postgres@localhost:5432/ontology?sslmode=disable" up

## Reset DB (CAUTION!)
reset-db:
	@echo "💣 Dropping and recreating DB..."
	docker exec -it $(DB_CONTAINER) dropdb -U postgres ontology || true
	docker exec -it $(DB_CONTAINER) createdb -U postgres ontology

## Run all tests
test:
	@echo "🧪 Running all tests..."
	go test ./tests/...

## Run tests with coverage
test-coverage:
	@echo "📊 Running tests with coverage..."
	go test -cover ./tests/...

## Run specific test categories
test-entities:
	@echo "🧪 Running entity tests..."
	go test ./tests/entities_test.go -v

test-causality:
	@echo "🧪 Running causality tests..."
	go test ./tests/causality_test.go -v

test-api:
	@echo "🧪 Running API integration tests..."
	go test ./tests/api_integration_test.go -v

test-philosophical:
	@echo "🧪 Running philosophical flow tests..."
	go test ./tests/philosophical_flow_test.go -v

## Run benchmarks
benchmark:
	@echo "⚡ Running benchmarks..."
	go test -bench=. ./tests/benchmark_test.go

## Run the test script
test-script:
	@echo "🚀 Running comprehensive test suite..."
	./run_tests.sh

## Help (show available commands)
help:
	@echo "Available commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  make %-15s %s\n", $$1, $$2}'
