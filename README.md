# Ontology-as-a-Service (OaaS)

**A backend platform in Go that models, queries, and reasons about entities, inspired by Neo-Aristotelian metaphysics (E. J. Lowe's four-category ontology).**

[![Go Version](https://img.shields.io/badge/Go-1.25-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![Tests](https://img.shields.io/badge/Tests-Passing-brightgreen.svg)](#testing)

---

## 🎯 Overview

OaaS implements **E. J. Lowe's Four-Category Ontology** as a modern knowledge graph backend, enabling:

- **Entity Modeling**: Substances, kinds, attributes, and modes as first-class objects
- **Causal Reasoning**: Aristotelian causes (material, formal, efficient, final)
- **Potentiality → Actuality**: Track and reason about what entities can become
- **Modern APIs**: REST, GraphQL, and real-time WebSocket support

## 📖 Philosophical Foundation

Based on **E. J. Lowe (1945–2014)**, a leading Neo-Aristotelian metaphysician, the system models reality through four categories:

1. **Substances** – Independent entities (e.g., a tree, a person)
2. **Kinds** – Natural classifications (e.g., oak, human)  
3. **Attributes** – General properties (e.g., color, weight)
4. **Modes** – Particular instantiations (e.g., *this tree's green leaf*)

The system integrates Aristotelian notions of **essence, potentiality/actuality, and the four causes**.

## 🚀 Quick Start

### Prerequisites
- Go 1.25+
- Docker & Docker Compose
- Make

### Installation

```bash
# Clone the repository
git clone https://github.com/apodicticscott/oaas.git
cd oaas

# Start services
make docker-up

# Run migrations
make migrate-up

# Start the server
make run
```

### Test the API

```bash
# Health check
curl http://localhost:8080/health

# Create a substance
curl -X POST http://localhost:8080/api/v1/substances \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Socrates",
    "kind": "Human",
    "essence": "Rational being with potentiality for wisdom"
  }'
```

## 🛠 Tech Stack

| Component | Technology |
|-----------|------------|
| **Language** | Go 1.25+ |
| **Web Framework** | Gin (REST) |
| **GraphQL** | gqlgen |
| **Database** | PostgreSQL |
| **ORM** | GORM |
| **Containerization** | Docker |
| **Testing** | testify, httptest |

## 📂 Project Structure

```
oaas/
├── cmd/server/           # Main application
├── internal/
│   ├── entities/         # Neo-Aristotelian entities
│   ├── causality/        # Potentiality → Actuality engine
│   ├── persistence/      # Database layer
│   └── api/             # REST/GraphQL handlers
├── graph/               # GraphQL schema & resolvers
├── db/migrations/       # Database migrations
├── tests/              # Comprehensive test suite
├── docker/             # Docker configuration
└── makefile           # Build & development commands
```

## 🧪 Testing

### Quick Testing

```bash
# Run all tests
make test

# Run with coverage
make test-coverage

# Run specific test categories
make test-entities      # Test Neo-Aristotelian entities
make test-causality     # Test potentiality → actuality
make test-philosophical # Test complete philosophical flow
make benchmark          # Performance benchmarks
```

### Test Categories

- **Unit Tests**: Individual entity and function testing
- **Integration Tests**: API endpoint testing  
- **Philosophical Tests**: Neo-Aristotelian ontology validation
- **Performance Tests**: Benchmark and load testing

## 📌 API Examples

### Complete Neo-Aristotelian Flow

```bash
# 1. Create a substance (independent entity)
curl -X POST http://localhost:8080/api/v1/substances \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Socrates",
    "kind": "Human",
    "essence": "Rational being with potentiality for wisdom"
  }'

# 2. Create a mode (particular instantiation)
curl -X POST http://localhost:8080/api/v1/modes \
  -H "Content-Type: application/json" \
  -d '{
    "value": "courageous",
    "substance_id": "SUBSTANCE_ID",
    "attribute_id": "attr-2"
  }'

# 3. Add the four Aristotelian causes
curl -X POST http://localhost:8080/api/v1/causes \
  -H "Content-Type: application/json" \
  -d '{
    "from_entity": "SUBSTANCE_ID",
    "to_entity": "flesh_bone_soul",
    "cause_type": "material"
  }'

# 4. Create a potentiality
curl -X POST http://localhost:8080/api/v1/potentialities \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Achieve Wisdom",
    "description": "Socrates can achieve philosophical wisdom",
    "conditions": "[{\"type\":\"mode\",\"name\":\"virtue\",\"value\":\"courageous\"}]",
    "substance_id": "SUBSTANCE_ID"
  }'

# 5. Actualize the potentiality
curl -X POST http://localhost:8080/api/v1/potentialities/POTENTIALITY_ID/actualize \
  -H "Content-Type: application/json" \
  -d '{
    "description": "Socrates achieved wisdom through philosophical inquiry"
  }'
```

### GraphQL

Visit **http://localhost:8080/playground** for interactive GraphQL exploration.

```graphql
query {
  substances {
    id
    name
    kind
    essence
    modes {
      value
      attribute { name }
    }
    potentialities {
      name
      description
    }
    actualities {
      description
      actualizedAt
    }
    causes {
      causeType
      toEntity
    }
  }
}
```

## 🔧 Development

### Available Commands

```bash
make help              # Show all available commands
make run               # Start the server
make build             # Build the binary
make docker-up         # Start Docker services
make docker-down       # Stop Docker services
make gqlgen            # Regenerate GraphQL code
make test              # Run all tests
make test-coverage     # Run tests with coverage
make benchmark         # Run performance tests
```

### Database Management

```bash
make psql              # Access PostgreSQL directly
make reset-db          # Reset database (CAUTION!)
```

## 🌐 API Endpoints

### REST API

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/health` | Health check |
| `GET` | `/api/v1/substances` | List all substances |
| `POST` | `/api/v1/substances` | Create substance |
| `GET` | `/api/v1/substances/:id` | Get substance by ID |
| `POST` | `/api/v1/modes` | Create mode |
| `POST` | `/api/v1/causes` | Add causal relation |
| `POST` | `/api/v1/potentialities` | Create potentiality |
| `POST` | `/api/v1/potentialities/:id/actualize` | Actualize potentiality |

### GraphQL

- **Playground**: `http://localhost:8080/playground`
- **Endpoint**: `http://localhost:8080/query`

## 🔬 Use Cases

- **Knowledge Graphs**: Scientific theories, research taxonomies
- **Philosophy-in-Tech**: Metaphysical concepts in backend design
- **Educational Tools**: Visualize Aristotle + Lowe's metaphysics
- **AI Reasoning**: Causal networks for ML/AI reasoning pipelines

## 📚 References

- E. J. Lowe, *The Four-Category Ontology: A Metaphysical Foundation for Natural Science* (2006)
- Kit Fine, *Essence and Modality*
- David Oderberg, *Real Essentialism*

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🧑‍💻 Resume Pitch

**Ontology-as-a-Service (OaaS)** – Backend system in **Go** implementing **E. J. Lowe's Neo-Aristotelian Four-Category Ontology**.
- Modeled substances, kinds, attributes, and modes with PostgreSQL
- Built a reasoning engine for potentiality → actuality transitions  
- Designed REST + GraphQL APIs for querying entities and Aristotelian causes
- Bridges **philosophy and distributed backend design** in a unique way

---

**Built with ❤️ and philosophical rigor**
