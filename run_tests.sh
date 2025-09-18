#!/bin/bash

# OaaS Test Runner Script
echo "🧪 Running OaaS Tests..."

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to run tests with output
run_test() {
    local test_name=$1
    local test_command=$2
    
    echo -e "${BLUE}Running: $test_name${NC}"
    echo "Command: $test_command"
    echo "----------------------------------------"
    
    if eval $test_command; then
        echo -e "${GREEN}✅ $test_name passed${NC}"
    else
        echo -e "${RED}❌ $test_name failed${NC}"
        return 1
    fi
    echo ""
}

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo -e "${RED}❌ Go is not installed or not in PATH${NC}"
    exit 1
fi

# Check if we're in the right directory
if [ ! -f "go.mod" ]; then
    echo -e "${RED}❌ Not in OaaS project directory. Please run from project root.${NC}"
    exit 1
fi

# Install test dependencies
echo -e "${YELLOW}Installing test dependencies...${NC}"
go mod tidy
go get github.com/stretchr/testify/assert
go get github.com/stretchr/testify/require

echo ""

# Run different types of tests
echo -e "${YELLOW}🧪 Running Unit Tests${NC}"
run_test "Entity Tests" "go test ./tests/entities_test.go -v"

echo -e "${YELLOW}🧪 Running Causality Engine Tests${NC}"
run_test "Causality Tests" "go test ./tests/causality_test.go -v"

echo -e "${YELLOW}🧪 Running API Integration Tests${NC}"
run_test "API Integration Tests" "go test ./tests/api_integration_test.go -v"

echo -e "${YELLOW}🧪 Running All Tests with Coverage${NC}"
run_test "All Tests with Coverage" "go test -cover ./tests/..."

echo -e "${YELLOW}🧪 Running Benchmark Tests${NC}"
run_test "Benchmark Tests" "go test -bench=. ./tests/benchmark_test.go"

echo -e "${YELLOW}🧪 Running Short Tests Only${NC}"
run_test "Short Tests" "go test -short ./tests/..."

echo ""
echo -e "${GREEN}🎯 Test Summary${NC}"
echo "========================================"

# Count test results
TOTAL_TESTS=$(go test -count=1 ./tests/... 2>&1 | grep -c "=== RUN")
PASSED_TESTS=$(go test -count=1 ./tests/... 2>&1 | grep -c "--- PASS")
FAILED_TESTS=$(go test -count=1 ./tests/... 2>&1 | grep -c "--- FAIL")

echo "Total Tests: $TOTAL_TESTS"
echo -e "Passed: ${GREEN}$PASSED_TESTS${NC}"
echo -e "Failed: ${RED}$FAILED_TESTS${NC}"

if [ $FAILED_TESTS -eq 0 ]; then
    echo -e "${GREEN}🎉 All tests passed!${NC}"
    exit 0
else
    echo -e "${RED}❌ Some tests failed. Check output above.${NC}"
    exit 1
fi
