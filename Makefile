.PHONY: help setup dev backend frontend test lint clean docker-up docker-down

help:
	@echo "TicksAndTomes Development Commands"
	@echo "=================================="
	@echo "setup            - Install all dependencies"
	@echo "dev              - Run backend + frontend in development mode"
	@echo "backend          - Run backend only"
	@echo "frontend         - Run frontend only"
	@echo "test             - Run all tests"
	@echo "test-backend     - Run backend tests"
	@echo "test-frontend    - Run frontend tests"
	@echo "lint             - Run linters"
	@echo "lint-backend     - Lint backend code"
	@echo "lint-frontend    - Lint frontend code"
	@echo "clean            - Clean build artifacts"
	@echo "docker-up        - Start Docker containers"
	@echo "docker-down      - Stop Docker containers"

setup:
	@echo "Installing dependencies..."
	@cd backend && go mod download
	@cd frontend && npm install

dev:
	@echo "Starting backend and frontend..."
	@make backend & make frontend

backend:
	@echo "Starting backend server..."
	@cd backend && go run ./cmd/server/main.go

frontend:
	@echo "Starting frontend dev server..."
	@cd frontend && npm start

test:
	@echo "Running all tests..."
	@make test-backend
	@make test-frontend

test-backend:
	@echo "Running backend tests..."
	@cd backend && go test -v ./...

test-frontend:
	@echo "Running frontend tests..."
	@cd frontend && npm test -- --coverage

lint:
	@echo "Running linters..."
	@make lint-backend
	@make lint-frontend

lint-backend:
	@echo "Linting backend..."
	@cd backend && golangci-lint run ./...

lint-frontend:
	@echo "Linting frontend..."
	@cd frontend && npm run lint

clean:
	@echo "Cleaning build artifacts..."
	@cd backend && go clean
	@cd frontend && rm -rf node_modules dist build
	@find . -name "*.o" -delete
	@find . -name "*.a" -delete

docker-up:
	@echo "Starting Docker containers..."
	@docker-compose up -d

docker-down:
	@echo "Stopping Docker containers..."
	@docker-compose down

db-init:
	@echo "Initializing database..."
	@cd backend && go run ./cmd/migrate/main.go
