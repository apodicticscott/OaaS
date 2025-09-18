#!/bin/bash

# OaaS Database Setup Script
echo "🚀 Setting up OaaS database..."

# Check if PostgreSQL is running
if ! pg_isready -h localhost -p 5433 -U postgres > /dev/null 2>&1; then
    echo "❌ PostgreSQL is not running. Please start it first:"
    echo "   docker-compose up -d"
    echo "   or start your local PostgreSQL instance"
    exit 1
fi

# Run migrations
echo "📝 Running database migrations..."

# Migration 1: Create core ontology tables
echo "  - Creating core ontology tables (four-category ontology)..."
psql -h localhost -p 5433 -U postgres -d ontology -f db/migrations/001_create_core_ontology_tables.up.sql

# Migration 2: Create causal relations
echo "  - Creating causal relations table (Aristotelian causes)..."
psql -h localhost -p 5433 -U postgres -d ontology -f db/migrations/002_create_causal_relations.up.sql

echo "✅ Database setup complete!"
echo ""
echo "🧪 Test the API with these commands:"
echo ""
echo "# Health check"
echo "curl http://localhost:8080/health"
echo ""
echo "# Get existing kinds"
echo "curl http://localhost:8080/api/v1/kinds"
echo ""
echo "# Get existing attributes"
echo "curl http://localhost:8080/api/v1/attributes"
echo ""
echo "# Create a substance"
echo "curl -X POST http://localhost:8080/api/v1/substances \\"
echo "  -H \"Content-Type: application/json\" \\"
echo "  -d '{\"name\": \"Tree-001\", \"kind\": \"Oak\", \"essence\": \"Living organism with potentiality for growth\"}'"
echo ""
echo "🎯 Start the server with: go run cmd/server/main.go"
