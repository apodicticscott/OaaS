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

### Health Check

```bash
# Check API health
curl -X GET http://localhost:8080/health
```

**Response:**
```json
{
  "status": "healthy",
  "message": "OaaS API is running"
}
```

### Core Entities

#### Substances (Independent Entities)

```bash
# Get all substances
curl -X GET http://localhost:8080/api/v1/substances

# Get specific substance by ID
curl -X GET http://localhost:8080/api/v1/substances/{substance_id}

# Create a new substance
curl -X POST http://localhost:8080/api/v1/substances \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Socrates",
    "kind": "Human", 
    "essence": "Rational being with potentiality for wisdom"
  }'

# Update a substance
curl -X PUT http://localhost:8080/api/v1/substances/{substance_id} \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Socrates the Wise",
    "kind": "Human",
    "essence": "Rational being who achieved wisdom"
  }'

# Delete a substance
curl -X DELETE http://localhost:8080/api/v1/substances/{substance_id}
```

**Sample Response:**
```json
{
  "id": "61ed27a3-7024-416c-a1ba-52165142dc1b",
  "name": "Socrates",
  "kind": "Human",
  "essence": "Rational being with potentiality for wisdom",
  "created_at": "2025-09-18T01:24:30.59444Z",
  "modes": [
    {
      "id": "79d383a2-6938-4502-8c1d-52a0e45e8ea9",
      "value": "courageous",
      "created_at": "2025-09-18T01:29:04.056296Z",
      "substance_id": "61ed27a3-7024-416c-a1ba-52165142dc1b",
      "attribute_id": "a024a9cb-73a9-44a4-ba14-97026d970bb7"
    }
  ]
}
```

#### Kinds (Natural Classifications)

```bash
# Get all kinds
curl -X GET http://localhost:8080/api/v1/kinds

# Create a new kind
curl -X POST http://localhost:8080/api/v1/kinds \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Dog",
    "description": "A domesticated canine animal"
  }'
```

**Sample Response:**
```json
{
  "kinds": [
    {
      "id": "2536ba77-2f15-4c8b-85fb-22065374ad7f",
      "name": "Dog",
      "description": "A domesticated canine animal",
      "created_at": "2025-09-18T01:30:55.926727663Z"
    }
  ]
}
```

#### Attributes (General Properties)

```bash
# Get all attributes
curl -X GET http://localhost:8080/api/v1/attributes

# Create a new attribute
curl -X POST http://localhost:8080/api/v1/attributes \
  -H "Content-Type: application/json" \
  -d '{
    "name": "height",
    "description": "Vertical measurement",
    "data_type": "number"
  }'
```

**Sample Response:**
```json
{
  "attributes": [
    {
      "id": "48d1e590-c4c7-42be-90f3-443a34cdb155",
      "name": "height",
      "description": "Vertical measurement",
      "data_type": "number",
      "created_at": "2025-09-18T01:30:53.840236937Z"
    }
  ]
}
```

#### Modes (Particular Instantiations)

```bash
# Get all modes
curl -X GET http://localhost:8080/api/v1/modes

# Create a new mode
curl -X POST http://localhost:8080/api/v1/modes \
  -H "Content-Type: application/json" \
  -d '{
    "value": "courageous",
    "substance_id": "61ed27a3-7024-416c-a1ba-52165142dc1b",
    "attribute_id": "a024a9cb-73a9-44a4-ba14-97026d970bb7"
  }'
```

**Sample Response:**
```json
{
  "modes": [
    {
      "id": "79d383a2-6938-4502-8c1d-52a0e45e8ea9",
      "value": "courageous",
      "created_at": "2025-09-18T01:29:04.056296Z",
      "substance_id": "61ed27a3-7024-416c-a1ba-52165142dc1b",
      "attribute_id": "a024a9cb-73a9-44a4-ba14-97026d970bb7",
      "substance": {
        "id": "61ed27a3-7024-416c-a1ba-52165142dc1b",
        "name": "Socrates",
        "kind": "Human",
        "essence": "Rational being with potentiality for wisdom",
        "created_at": "2025-09-18T01:24:30.59444Z"
      },
      "attribute": {
        "id": "a024a9cb-73a9-44a4-ba14-97026d970bb7",
        "name": "virtue",
        "description": "Moral excellence and character",
        "data_type": "string",
        "created_at": "2025-09-18T01:28:59.961002Z"
      }
    }
  ]
}
```

### Causality & Potentiality

#### Aristotelian Causes

```bash
# Add a causal relation
curl -X POST http://localhost:8080/api/v1/causes \
  -H "Content-Type: application/json" \
  -d '{
    "from_entity": "61ed27a3-7024-416c-a1ba-52165142dc1b",
    "to_entity": "philosophical_inquiry",
    "cause_type": "efficient"
  }'

# Get causes for a substance
curl -X GET http://localhost:8080/api/v1/substances/{substance_id}/causes
```

**Sample Response:**
```json
{
  "efficient": "philosophical_inquiry"
}
```

#### Potentialities & Actualities

```bash
# Get all potentialities
curl -X GET http://localhost:8080/api/v1/potentialities

# Create a potentiality
curl -X POST http://localhost:8080/api/v1/potentialities \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Achieve Wisdom",
    "description": "Socrates can achieve philosophical wisdom through inquiry",
    "conditions": "[{\"type\":\"mode\",\"name\":\"virtue\",\"value\":\"courageous\"}]",
    "substance_id": "61ed27a3-7024-416c-a1ba-52165142dc1b"
  }'

# Check if potentiality can be actualized
curl -X GET http://localhost:8080/api/v1/potentialities/{potentiality_id}/conditions

# Actualize a potentiality
curl -X POST http://localhost:8080/api/v1/potentialities/{potentiality_id}/actualize \
  -H "Content-Type: application/json" \
  -d '{
    "description": "Socrates achieved wisdom through philosophical inquiry and courageous questioning"
  }'

# Get substance evolution
curl -X GET http://localhost:8080/api/v1/substances/{substance_id}/evolution
```

**Sample Responses:**

*Potentiality Creation:*
```json
{
  "id": "c461b7ac-bb5c-4d08-97bf-419207ac28ee",
  "name": "Achieve Wisdom",
  "description": "Socrates can achieve philosophical wisdom through inquiry",
  "conditions": "[{\"type\":\"mode\",\"name\":\"virtue\",\"value\":\"courageous\"}]",
  "created_at": "2025-09-18T01:31:12.765685Z",
  "substance_id": "61ed27a3-7024-416c-a1ba-52165142dc1b"
}
```

*Condition Check:*
```json
{
  "can_actualize": true,
  "unmet_conditions": null
}
```

*Actuality Creation:*
```json
{
  "id": "b81aff87-28fb-4d97-b772-20b850d9455c",
  "description": "Socrates achieved wisdom through philosophical inquiry and courageous questioning",
  "actualized_at": "2025-09-18T01:33:18.636639431Z",
  "substance_id": "61ed27a3-7024-416c-a1ba-52165142dc1b",
  "potentiality_id": "c461b7ac-bb5c-4d08-97bf-419207ac28ee"
}
```

### Complete Neo-Aristotelian Flow Example

Here's a complete example demonstrating the full philosophical flow:

```bash
# 1. Create a substance (independent entity)
SUBSTANCE_RESPONSE=$(curl -s -X POST http://localhost:8080/api/v1/substances \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Socrates",
    "kind": "Human",
    "essence": "Rational being with potentiality for wisdom"
  }')

SUBSTANCE_ID=$(echo $SUBSTANCE_RESPONSE | jq -r '.id')

# 2. Create an attribute (general property)
ATTRIBUTE_RESPONSE=$(curl -s -X POST http://localhost:8080/api/v1/attributes \
  -H "Content-Type: application/json" \
  -d '{
    "name": "virtue",
    "description": "Moral excellence and character",
    "data_type": "string"
  }')

ATTRIBUTE_ID=$(echo $ATTRIBUTE_RESPONSE | jq -r '.id')

# 3. Create a mode (particular instantiation)
curl -X POST http://localhost:8080/api/v1/modes \
  -H "Content-Type: application/json" \
  -d "{
    \"value\": \"courageous\",
    \"substance_id\": \"$SUBSTANCE_ID\",
    \"attribute_id\": \"$ATTRIBUTE_ID\"
  }"

# 4. Add Aristotelian causes
curl -X POST http://localhost:8080/api/v1/causes \
  -H "Content-Type: application/json" \
  -d "{
    \"from_entity\": \"$SUBSTANCE_ID\",
    \"to_entity\": \"philosophical_inquiry\",
    \"cause_type\": \"efficient\"
  }"

# 5. Create a potentiality
POTENTIALITY_RESPONSE=$(curl -s -X POST http://localhost:8080/api/v1/potentialities \
  -H "Content-Type: application/json" \
  -d "{
    \"name\": \"Achieve Wisdom\",
    \"description\": \"Socrates can achieve philosophical wisdom\",
    \"conditions\": \"[{\\\"type\\\":\\\"mode\\\",\\\"name\\\":\\\"virtue\\\",\\\"value\\\":\\\"courageous\\\"}]\",
    \"substance_id\": \"$SUBSTANCE_ID\"
  }")

POTENTIALITY_ID=$(echo $POTENTIALITY_RESPONSE | jq -r '.id')

# 6. Check if conditions are met
curl -X GET http://localhost:8080/api/v1/potentialities/$POTENTIALITY_ID/conditions

# 7. Actualize the potentiality
curl -X POST http://localhost:8080/api/v1/potentialities/$POTENTIALITY_ID/actualize \
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
| **Substances** | | |
| `GET` | `/api/v1/substances` | List all substances |
| `GET` | `/api/v1/substances/:id` | Get substance by ID |
| `POST` | `/api/v1/substances` | Create substance |
| `PUT` | `/api/v1/substances/:id` | Update substance |
| `DELETE` | `/api/v1/substances/:id` | Delete substance |
| **Kinds** | | |
| `GET` | `/api/v1/kinds` | List all kinds |
| `POST` | `/api/v1/kinds` | Create kind |
| **Attributes** | | |
| `GET` | `/api/v1/attributes` | List all attributes |
| `POST` | `/api/v1/attributes` | Create attribute |
| **Modes** | | |
| `GET` | `/api/v1/modes` | List all modes |
| `POST` | `/api/v1/modes` | Create mode |
| **Causality** | | |
| `GET` | `/api/v1/substances/:id/causes` | Get causes for substance |
| `POST` | `/api/v1/causes` | Add causal relation |
| **Potentialities** | | |
| `GET` | `/api/v1/potentialities` | List all potentialities |
| `POST` | `/api/v1/potentialities` | Create potentiality |
| `GET` | `/api/v1/potentialities/:id/conditions` | Check potentiality conditions |
| `POST` | `/api/v1/potentialities/:id/actualize` | Actualize potentiality |
| **Evolution** | | |
| `GET` | `/api/v1/substances/:id/evolution` | Get substance evolution |

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
