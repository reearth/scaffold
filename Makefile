TEST_DIR ?= ./...

SCHEMATYPER := github.com/idubinskiy/schematyper
MANIFEST_DIR := pkg/plugin/manifest

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
	@echo "  dev-install       Install tools for make dev - behavior not confirmed in windows."
	@echo "  dev               Run the application with hot reloading"
	@echo "  run-app           Run the application"
	@echo "  run-db            Run the MongoDB database using Docker Compose"
	@echo "  up-gcs            Run the fake-gcs-server using Docker Compose"
	@echo "  down-gcs          Stop the fake-gcs-server using Docker Compose"

lint:
	golangci-lint run --fix

TARGET_TEST :=./...
REEARTH_DB := mongodb://localhost
test:
	REEARTH_DB=${REEARTH_DB} go test ./... -run ${TARGET_TEST}

failcheck:
	go test -race -short -failfast -p 1 $(TEST_DIR)

AIR_BIN := $(shell which air)
dev-install:
ifndef AIR_BIN
	@echo "reflex is not installed. Installing..."
	@go install github.com/air-verse/air@v1.61.5
else
	@echo "air is already installed."
endif

dev: dev-install
	air

build:
	go build ./cmd/app

run-app:
	go run ./cmd/app

run-db:
	docker compose -f ./compose.yml up mongo

generate: dev-install
	go generate ./...

run-gcs:
	docker compose -f ./compose.yml up gcs

down-gcs:
	docker compose -f ./compose.yml down gcs

.PHONY: lint test failcheck e2e build dev-install dev run-app run-db gql up-gcs down-gcs mockuser schematyper