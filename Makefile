# Constants
PROJECT_NAME := txgo
MIGRATIONS_DIR := ./db/migration
DATABASE_HOST ?= localhost
DATABASE_PORT ?= 3306
DATABASE_USER ?= user
DATABASE_PASSWORD ?= password
DATABASE_NAME = db
DATABASE_URL := "mysql://${DATABASE_USER}:${DATABASE_PASSWORD}@tcp(${DATABASE_HOST}:${DATABASE_PORT})/${DATABASE_NAME}"
SERVER_PORT := 3000

# Build
.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux go build -o ./bin/${PROJECT_NAME} ./cmd/api/main.go

# Migrations
.PHONY: mig-up
mig-up: ## Runs the migrations up
	migrate -path ${MIGRATIONS_DIR} -database ${DATABASE_URL} up

.PHONY: mig-down
mig-down: ## Runs the migrations down
	migrate -path ${MIGRATIONS_DIR} -database ${DATABASE_URL} down

.PHONY: new-mig
new-mig:
	migrate create -ext sql -dir ${MIGRATIONS_DIR} -seq $(NAME)
