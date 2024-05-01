include .env

MAIN_PACKAGE_PATH := ./cmd/ordersystem
BINARY_NAME := ordersystem


# ============================================================================ #
# HELPERS
# ============================================================================ #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'


# ============================================================================ #
# Database Migration
# ============================================================================ #
DATABASE_URL="$(DB_DRIVER)://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)"

## migrate: Apply all peding migrations to the database
.PHONY: migrate
migrate:
	migrate -path internal/infra/database/sql/migrations -database $(DATABASE_URL) up

## migrate/down: Revert migrations in the database
migrate/down:
	migrate -path internal/infra/database/sql/migrations -database $(DATABASE_URL) down

## migrate/create: Create a new migration use `make migrate/create name=<MIGRATION_NAME>`
.PHONY: migrate/create
migrate/create:
	migrate create -ext sql -dir internal/infra/database/sql/migrations $(name)


# ============================================================================ #
# Code Generation
# ============================================================================ #

## generate/grpc: Generate gRPC/protoc code
.PHONY: generate/grpc
generate/grpc:
	protoc --go_out=. --go-grpc_out=. internal/infra/grpc/protofiles/order.proto

## generate/graphql: Generate GraphQL code
.PHONY: generate/graphql
generate/graphql:
	go run github.com/99designs/gqlgen generate


# ============================================================================ #
# Development
# ============================================================================ #

## build: build the application
.PHONY: build
build:
	go build -o=/tmp/bin/${BINARY_NAME} ${MAIN_PACKAGE_PATH}

## run: run the  application
.PHONY: run
run: build
	/tmp/bin/$(BINARY_NAME)

## run/live: run the application with reloading on file changes
.PHONY: run/live
run/live:
	go run github.com/cosmtrek/air@v1.51.0 \
			--build.cmd "make build" --build.bin "/tmp/bin/${BINARY_NAME}" --build.delay "100" \
			--build.exclude_dir "" \
			--build.include_ext "go" \
			--misc.clean_on_exit "true"