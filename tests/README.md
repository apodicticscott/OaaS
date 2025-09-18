# OaaS Testing Suite

This directory contains comprehensive tests for the Ontology-as-a-Service (OaaS) system, implementing E.J. Lowe's Neo-Aristotelian Four-Category Ontology.

## 🧪 Test Structure

### Unit Tests
- **`entities_test.go`** - Tests for core entity creation and validation
- **`causality_test.go`** - Tests for the causality engine and potentiality → actuality transitions
- **`api_integration_test.go`** - Integration tests for REST API endpoints

### Performance Tests
- **`benchmark_test.go`** - Benchmark tests for performance analysis

### Configuration
- **`test_config.go`** - Test configuration and utilities

## 🚀 Running Tests

### Quick Start
```bash
# Run all tests
./run_tests.sh

# Or run specific test types
go test ./tests/...
go test -cover ./tests/...
go test -bench=. ./tests/...
```

### Individual Test Categories

#### Entity Tests
```bash
go test ./tests/entities_test.go -v
```
Tests the core Neo-Aristotelian entities:
- Substances (independent entities)
- Kinds (natural classifications)
- Attributes (general properties)
- Modes (particular instantiations)
- Causal Relations (Aristotelian causes)
- Potentialities (what substances can become)
- Actualities (realized potentialities)

#### Causality Engine Tests
```bash
go test ./tests/causality_test.go -v
```
Tests the philosophical reasoning engine:
- Condition checking for potentialities
- Potentiality → actuality transitions
- Four Aristotelian causes (material, formal, efficient, final)
- Substance evolution tracking

#### API Integration Tests
```bash
go test ./tests/api_integration_test.go -v
```
Tests the REST API endpoints:
- CRUD operations for all entities
- Error handling and validation
- HTTP status codes and responses
- Data serialization/deserialization

#### Benchmark Tests
```bash
go test -bench=. ./tests/benchmark_test.go
```
Performance tests for:
- Entity creation speed
- Database query performance
- Condition checking efficiency
- Actualization process speed

## 📊 Test Coverage

Run tests with coverage analysis:
```bash
go test -cover ./tests/...
go test -coverprofile=coverage.out ./tests/...
go tool cover -html=coverage.out
```

## 🎯 Test Philosophy

The tests are designed to validate both the **technical implementation** and the **philosophical correctness** of the Neo-Aristotelian ontology:

### Technical Validation
- ✅ Data integrity and relationships
- ✅ API contract compliance
- ✅ Error handling and edge cases
- ✅ Performance and scalability

### Philosophical Validation
- ✅ Four-category ontology structure
- ✅ Proper causal relationships
- ✅ Potentiality → actuality transitions
- ✅ Aristotelian cause types

## 🔧 Test Configuration

### Environment Variables
- `SKIP_DB_TESTS=true` - Skip database-dependent tests
- `TEST_DB_URL` - Custom test database URL

### Test Database
Tests use in-memory SQLite by default for speed and isolation. For integration testing with PostgreSQL:

```bash
# Set up test database
createdb oaas_test

# Run tests with PostgreSQL
TEST_DB_URL="postgres://user:pass@localhost/oaas_test" go test ./tests/...
```

## 📝 Writing New Tests

### Test Naming Convention
- `TestFunctionName` - Unit tests
- `TestFunctionName_Scenario` - Specific scenario tests
- `BenchmarkFunctionName` - Performance tests

### Test Structure
```go
func TestExample(t *testing.T) {
    // Arrange
    setup := setupTest(t)
    defer cleanup(setup)
    
    // Act
    result := functionUnderTest()
    
    // Assert
    assert.Equal(t, expected, result)
}
```

### Best Practices
1. **Isolation** - Each test should be independent
2. **Clarity** - Test names should describe what's being tested
3. **Coverage** - Test both happy path and edge cases
4. **Philosophy** - Validate ontological correctness, not just technical correctness

## 🐛 Debugging Tests

### Verbose Output
```bash
go test -v ./tests/...
```

### Specific Test
```bash
go test -run TestSpecificFunction ./tests/...
```

### Race Detection
```bash
go test -race ./tests/...
```

### Memory Profiling
```bash
go test -memprofile=mem.prof ./tests/...
go tool pprof mem.prof
```

## 📈 Continuous Integration

The test suite is designed to run in CI/CD pipelines:

```yaml
# Example GitHub Actions
- name: Run Tests
  run: |
    go test -v ./tests/...
    go test -cover ./tests/...
    go test -bench=. ./tests/...
```

## 🎓 Educational Value

These tests serve as **living documentation** of the Neo-Aristotelian ontology implementation. They demonstrate:

- How to model philosophical concepts in code
- How to implement causal reasoning systems
- How to build APIs that respect ontological principles
- How to test both technical and philosophical correctness

## 🔗 Related Documentation

- [Main README](../README.md) - Project overview and philosophy
- [API Documentation](../README.md#testing) - REST API testing examples
- [GraphQL Schema](../graph/schema.graphqls) - GraphQL testing examples
