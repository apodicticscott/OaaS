#!/bin/bash

# OaaS API Test Script
echo "🧪 Testing OaaS API..."

# Base URL
BASE_URL="http://localhost:8080"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Function to test API endpoint
test_endpoint() {
    local method=$1
    local endpoint=$2
    local data=$3
    local description=$4
    
    echo -e "${YELLOW}Testing: $description${NC}"
    
    if [ -n "$data" ]; then
        response=$(curl -s -X $method "$BASE_URL$endpoint" \
            -H "Content-Type: application/json" \
            -d "$data")
    else
        response=$(curl -s -X $method "$BASE_URL$endpoint")
    fi
    
    if [ $? -eq 0 ]; then
        echo -e "${GREEN}✅ Success${NC}"
        echo "$response" | jq . 2>/dev/null || echo "$response"
    else
        echo -e "${RED}❌ Failed${NC}"
        echo "$response"
    fi
    echo ""
}

# Test 1: Health check
test_endpoint "GET" "/health" "" "Health Check"

# Test 2: Get existing kinds
test_endpoint "GET" "/api/v1/kinds" "" "Get Kinds"

# Test 3: Get existing attributes
test_endpoint "GET" "/api/v1/attributes" "" "Get Attributes"

# Test 4: Create a substance
test_endpoint "POST" "/api/v1/substances" '{
    "name": "Tree-001",
    "kind": "Oak",
    "essence": "Living organism with potentiality for growth"
}' "Create Substance"

# Test 5: Get all substances
test_endpoint "GET" "/api/v1/substances" "" "Get All Substances"

# Test 6: Create a mode (you'll need to get IDs from previous responses)
echo -e "${YELLOW}Note: For modes, you'll need to get substance and attribute IDs from the responses above${NC}"
echo ""

# Test 7: Add a causal relation
test_endpoint "POST" "/api/v1/causes" '{
    "from_entity": "test-substance",
    "to_entity": "wood_water_carbon",
    "cause_type": "material"
}' "Add Material Cause"

# Test 8: Create a potentiality
test_endpoint "POST" "/api/v1/potentialities" '{
    "name": "Grow Leaves",
    "description": "The tree can grow new leaves in spring",
    "conditions": "[{\"type\":\"attribute\",\"name\":\"season\",\"value\":\"spring\"}]",
    "substance_id": "test-substance"
}' "Create Potentiality"

echo -e "${GREEN}🎯 Test complete! Check the responses above for any errors.${NC}"
echo ""
echo "Next steps:"
echo "1. Start the server: go run cmd/server/main.go"
echo "2. Run this test: ./test_api.sh"
echo "3. Visit GraphQL playground: http://localhost:8080/playground"
