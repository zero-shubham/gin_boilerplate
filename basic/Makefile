#!make
include .env

project:=basic

ARCH := amd64

# Database Operations
.PHONY: migrate
up:
	migrate -database ${MIGRATE_DB_URI} -path ./migrations up

down:
	migrate -database ${MIGRATE_DB_URI} -path ./migrations down

new:
	migrate create -ext sql -dir migrations -seq ${filename}

# Generate cli
.PHONY: cli
cli:
	go build -o cmd cli/main.go

.PHONY: database
start:
	docker-compose up postgres-service