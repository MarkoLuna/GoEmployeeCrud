# Employee CRUD API

A RESTful API for employee management built with Go, PostgreSQL, and Docker.

## Overview

This project provides a complete CRUD (Create, Read, Update, Delete) API for managing employee records. It includes authentication, validation, and supports multiple deployment options including local development, Docker, and Kubernetes.

## Features

- RESTful API design
- Employee management with full CRUD operations
- PostgreSQL database integration
- JWT authentication
- Input validation
- Docker containerization
- Kubernetes deployment support
- Health check endpoints
- Comprehensive test coverage

## Requirements

- Go 1.25+
- PostgreSQL
- Docker (optional)
- Kubernetes (optional)

## Installation

### Clone the repository
```bash
git clone https://github.com/MarkoLuna/GoEmployeeCrud.git
cd GoEmployeeCrud
```

### Install dependencies
```bash
go mod download
```

### Environment setup
Copy the example environment file and configure your database settings:
```bash
cp .env.example .env
# Edit .env with your database configuration
```

## Running the Application

### Local Development
```bash
# Run directly
go run cmd/server/main.go

# Or using Make
make run
```

### Docker
```bash
# Build and run with Docker
make docker-run

# Or using Docker Compose
make docker-compose-run
```

### Kubernetes (Local)
```bash
# Deploy to local Kubernetes cluster
make k8-apply

# Test the deployment
kubectl exec -it employeecrud-pod -- apk add curl
kubectl exec -it employeecrud-pod -- curl -X GET http://employeecrud-service:8080/healthcheck/

# Clean up
make k8-remove
```

## API Endpoints

### Health Check
```bash
curl -X GET http://localhost:8080/healthcheck/
# Response: OK
```

### Employee Management

#### Get All Employees
```bash
curl http://localhost:8080/api/employee/
```

#### Get Employee by ID
```bash
curl http://localhost:8080/api/employee/{id}
```

#### Create Employee
```bash
curl --location --request POST 'http://localhost:8080/api/employee/' \
--header 'Content-Type: application/json' \
--data-raw '{
    "firstName": "Marcos",
    "lastName": "Luna",
    "secondLastName": "Valdez",
    "dateOfBirth": "1994-04-25T12:00:00Z",
    "dateOfEmployment": "1994-04-25T12:00:00Z",
    "status": "ACTIVE"
}'
```

#### Update Employee
```bash
curl --location --request PUT 'http://localhost:8080/api/employee/{id}' \
--header 'Content-Type: application/json' \
--data-raw '{
    "firstName": "Gerardo",
    "lastName": "Luna",
    "secondLastName": "Valdezz",
    "dateOfBirth": "1994-04-25T12:00:00Z",
    "dateOfEmployment": "0001-01-01T00:00:00Z",
    "status": "INACTIVE"
}'
```

#### Delete Employee
```bash
curl -X DELETE 'http://localhost:8080/api/employee/{id}'
```

## Testing

Run the test suite:
```bash
go test ./...
```

### Available Make Commands

The project includes a Makefile with the following commands:

#### Build and Run
```bash
make build          # Build the application binary
make run            # Build and run the application locally
make clean          # Clean build artifacts and binaries
```

#### Testing
```bash
make test           # Run tests with race detection
make test-cover     # Run tests with coverage information
make test-total-cover # Run tests and display total coverage percentage
make vet            # Run go vet for static analysis
```

#### Docker
```bash
make docker-build   # Build Docker image for Linux/amd64
make docker-run     # Build and run Docker container
make docker-compose-run # Build and run with Docker Compose
```

#### Kubernetes
```bash
make k8-apply       # Deploy to Kubernetes cluster
make k8-remove      # Remove Kubernetes resources
```

#### Examples
```bash
# Development workflow
make build && make run

# Testing workflow
make test && make vet && make test-total-cover

# Docker workflow
make docker-build && make docker-run

# Kubernetes deployment
make k8-apply
# Test the deployment
make k8-remove
```

## Project Structure

```
├── cmd/
│   └── server/
│       └── main.go          # Application entry point
├── pkg/                     # Application packages
│   ├── handlers/            # HTTP handlers
│   ├── models/              # Data models
│   ├── services/            # Business logic
│   └── utils/               # Utilities
├── k8s/                     # Kubernetes manifests
├── resources/               # Configuration files
├── docker-compose.yml       # Docker Compose configuration
├── Dockerfile              # Docker image configuration
├── Makefile                # Build and deployment scripts
└── go.mod                  # Go module dependencies
```

## Technology Stack

- **Language**: Go 1.25+
- **Web Framework**: Gorilla Mux
- **Database**: PostgreSQL
- **Authentication**: JWT
- **Validation**: go-playground/validator
- **Testing**: Testify + SQL Mock
- **Containerization**: Docker
- **Orchestration**: Kubernetes

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for new functionality
5. Run the test suite
6. Submit a pull request

## License

This project is licensed under the MIT License.
