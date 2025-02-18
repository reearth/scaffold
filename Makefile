TEST_DIR ?= ./...

default: help

help:
	@echo "Usage:"
	@echo "  make <target>"
	@echo ""
	@echo "Targets:"
	@echo "  lint              Run golangci-lint with auto-fix"
	@echo "  test              Run unit tests with race detector in short mode"
	@echo "  failcheck         Run unit tests with fail-fast and no parallel execution"
	@echo "  build             Build the project"
	@echo "  dev               Run the application with hot reloading"
	@echo "  generate          Run go generate"
	@echo "  run-app           Run the application"
	@echo "  run-db            Run the MongoDB database using Docker Compose"
	@echo "  down-db           Stop the MongoDB database using Docker Compose"
	@echo "  up-gcs            Run the fake-gcs-server using Docker Compose"
	@echo "  down-gcs          Stop the fake-gcs-server using Docker Compose"

lint:
	golangci-lint run --fix

TARGET_TEST ?=
REEARTH_DB := mongodb://localhost
test:
	@if [ -z "$(TARGET_TEST)" ]; then \
		REEARTH_DB=$(REEARTH_DB) go test ./... ; \
	else \
		REEARTH_DB=$(REEARTH_DB) go test ./... -run "$(TARGET_TEST)"; \
	fi

failcheck:
	go test -race -short -failfast -p 1 $(TEST_DIR)

dev:
	go run github.com/air-verse/air

build:
	go build ./cmd/app

generate:
	go generate ./...

run-app:
	go run ./cmd/app

run-db:
	docker compose -f ./compose.yml up mongo

down-db:
	docker compose -f ./compose.yml down mongo

run-gcs:
	docker compose -f ./compose.yml up gcs

down-gcs:
	docker compose -f ./compose.yml down gcs

.PHONY: lint test failcheck build dev generate run-app run-db down-db run-gcs down-gcs
