SHELL := bash

# Define OS
ifeq ($(OS),Windows_NT)
    MIGRATE_ACTION = scripts\bat\migrate-action.bat
    MIGRATE_CREATE = scripts\bat\migrate-create.bat
    RM = del /Q
else
    MIGRATE_ACTION = ./scripts/sh/migrate-action.sh
    MIGRATE_CREATE = ./scripts/sh/migrate-create.sh
    RM = rm -f
endif

# Load .env
include .env
export

# ========= Commands =========

# Up environment docker containers
env-up:
	docker compose up -d affiliate-system-postgres
	docker compose up -d affiliate-system-kafka
	docker compose up -d affiliate-system-wal-listener

# Down environment docker containers
env-down:
	docker compose down -v affiliate-system-wal-listener
	docker compose down -v affiliate-system-kafka
	docker compose down -v affiliate-system-postgres

# Create db migration
migrate-create:
	$(MIGRATE_CREATE) $(seq)

# Up db migration
migrate-up:
	$(MIGRATE_ACTION) up

# Down db migration
migrate-down:
	$(MIGRATE_ACTION) down

# Use db migration with custom action
migrate-action:
	$(MIGRATE_ACTION) $(action)

# Check migrate status
migrate-status:
	$(MIGRATE_ACTION) status

# Run http server
run-server:
	@cd backend && go mod tidy && go run cmd/server/main.go

# Run kafka consumer
run-consumer:
	@cd backend && go mod tidy && go run cmd/consumer/main.go

.PHONY: postgres-up postgres-down kafka-up kafka-down migrate-create migrate-up migrate-down migrate-action run-server run-consumer migrate-status migrate-force