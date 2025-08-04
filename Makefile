# Database configuration
DB_HOST ?= localhost
DB_PORT ?= 5432
DB_NAME ?= techschool
DB_USER ?= postgres
DB_PASSWORD ?= postgres

# Flyway commands
.PHONY: migrate migrate-info migrate-clean migrate-repair migrate-validate

# Run all pending migrations
migrate:
	flyway -configFiles=flyway.conf migrate

# Show migration status
migrate-info:
	flyway -configFiles=flyway.conf info

# Clean database (drop all tables)
migrate-clean:
	flyway -configFiles=flyway.conf clean

# Repair migration history
migrate-repair:
	flyway -configFiles=flyway.conf repair

# Validate migrations
migrate-validate:
	flyway -configFiles=flyway.conf validate

# Create new migration file
migrate-create:
	@read -p "Enter migration description: " desc; \
	flyway -configFiles=flyway.conf -locations=filesystem:migrations migrate:create -name="$$desc"

# Build and run application
build:
	go build -o bin/app main.go

run: build
	./bin/app

# Docker commands
docker-up:
	@echo "Starting services with Docker Compose..."
	@docker-compose up -d

docker-down:
	@echo "Stopping services..."
	@docker-compose down

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf bin/
	@go clean

# Install dependencies
deps:
	go mod tidy
	go mod download

# Test
test:
	go test ./...

# Docker commands
docker-build:
	docker build -t techschool-app .

docker-run:
	docker run -p 8080:8080 techschool-app

# Database setup
db-setup:
	@echo "Setting up database..."
	@echo "Make sure PostgreSQL is running and database '$(DB_NAME)' exists"
	@echo "You can create it with: createdb $(DB_NAME)"

# Help
help:
	@echo "Available commands:"
	@echo "  migrate          - Run all pending migrations"
	@echo "  migrate-info     - Show migration status"
	@echo "  migrate-clean    - Clean database (drop all tables)"
	@echo "  migrate-repair   - Repair migration history"
	@echo "  migrate-validate - Validate migrations"
	@echo "  migrate-create   - Create new migration file"
	@echo "  build           - Build application"
	@echo "  run             - Build and run application"
	@echo "  dev             - Run in development mode"
	@echo "  clean           - Clean build artifacts"
	@echo "  deps            - Install dependencies"
	@echo "  test            - Run tests"
	@echo "  docker-build    - Build Docker image"
	@echo "  docker-run      - Run Docker container"
	@echo "  db-setup        - Database setup instructions" 