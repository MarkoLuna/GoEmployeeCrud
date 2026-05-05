# =============================================================================
# Employee CRUD Application Makefile
# =============================================================================
# This Makefile provides commands for building, testing, and deploying the
# Employee CRUD REST API application.
#
# Usage:
#   make <target>    # Execute the specified target
#   make help        # Show available targets (if implemented)
# =============================================================================

# Constants
NAME=employeeCrudApp                                    # Application binary name
PROJECT?=github.com/MarkoLuna/GoEmployeeCrud

# =============================================================================
# BUILD AND RUN TARGETS
# =============================================================================

.PHONY: build
build:                                                  # Build the application binary for current platform
	@echo "Building application..."
	go build -o ${NAME} "${PROJECT}/cmd/server"
	@echo "Build complete: ${NAME}"

.PHONY: run
run: build                                              # Build and run the application locally
	@echo "Starting application..."
	./${NAME}

.PHONY: clean
clean:                                                  # Clean build artifacts and binaries
	@echo "Cleaning up..."
	go clean "${PROJECT}/..."
	rm -f ${NAME}
	@echo "Clean complete"

# =============================================================================
# TESTING TARGETS
# =============================================================================

.PHONY: test
test:                                                   # Run tests with race detection enabled
	@echo "Running tests with race detection..."
	go test -race "${PROJECT}/..."

.PHONY: test-cover
test-cover:                                             # Run tests with coverage information
	@echo "Running tests with coverage..."
	go test -cover "${PROJECT}/..."

.PHONY: vet
vet:                                                    # Run go vet for static analysis
	@echo "Running static analysis..."
	go vet "${PROJECT}/..."

.PHONY: test-total-cover
test-total-cover:                                       # Run tests and display total coverage percentage
	@echo "Running tests with coverage profile..."
	go test "${PROJECT}/..." -coverprofile cover.out > /dev/null
	@echo "Coverage report:"
	go tool cover -func cover.out
	rm cover.out

# =============================================================================
# DOCKER TARGETS
# =============================================================================

.PHONY: docker-build
docker-build:                                           # Build Docker image for Linux/amd64 architecture
	@echo "Building Docker image..."
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${NAME} "${PROJECT}/cmd/server"
	docker build -t goemployee_crud:latest .
	rm -f ${NAME}
	@echo "Docker image built: goemployee_crud:latest"

.PHONY: docker-run
docker-run: docker-build                                # Build and run Docker container
	@echo "Running Docker container..."
	docker run -it -p 8080:8080 --rm goemployee_crud

.PHONY: docker-compose-run
docker-compose-run: docker-build                         # Build and run with Docker Compose
	@echo "Starting with Docker Compose..."
	docker-compose up

# =============================================================================
# KUBERNETES TARGETS
# =============================================================================

.PHONY: k8-apply
k8-apply: docker-build                                  # Deploy application to Kubernetes cluster
	@echo "Deploying to Kubernetes..."
	kubectl apply -f k8s/pod.yaml
	kubectl apply -f k8s/service.yaml
	@echo "Kubernetes deployment complete"

.PHONY: k8-remove
k8-remove:                                              # Remove Kubernetes resources
	@echo "Removing Kubernetes resources..."
	kubectl delete pod employeecrud-pod || true
	kubectl delete service employeecrud-service || true
	@echo "Kubernetes resources removed"

# =============================================================================
# HELP TARGET (Optional)
# =============================================================================

.PHONY: help
help:                                                   # Show available targets with descriptions
	@echo "Available targets:"
	@echo "  build              - Build the application binary"
	@echo "  run                - Build and run the application locally"
	@echo "  clean              - Clean build artifacts and binaries"
	@echo "  test               - Run tests with race detection"
	@echo "  test-cover         - Run tests with coverage information"
	@echo "  test-total-cover   - Run tests and display total coverage percentage"
	@echo "  vet                - Run go vet for static analysis"
	@echo "  docker-build       - Build Docker image for Linux/amd64"
	@echo "  docker-run         - Build and run Docker container"
	@echo "  docker-compose-run - Build and run with Docker Compose"
	@echo "  k8-apply           - Deploy to Kubernetes cluster"
	@echo "  k8-remove          - Remove Kubernetes resources"
	@echo "  help               - Show this help message"
