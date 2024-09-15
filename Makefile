include .envrc

# ----------------------------------------------------------------------------------------
# Proto
# ----------------------------------------------------------------------------------------
.PHONY: proto_build
proto_build:
	protoc --go_out=./internal/proto --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=./internal/proto ./internal/proto/*.proto


# ----------------------------------------------------------------------------------------
# Runs
# ----------------------------------------------------------------------------------------
.PHONY: run/rpc
run/rpc:
	go run ./cmd/rpc -db-dsn=${DB_DSN}

.PHONY: run/rest
run/rest:
	go run ./cmd/rest -db-dsn=${DB_DSN}

.PHONY: run/web
run/web:
	go run ./cmd/web -db-dsn=${DB_DSN}


# ----------------------------------------------------------------------------------------
# Migration
# ----------------------------------------------------------------------------------------
.PHONY: migration/new
migration/new:
	@echo "Creating migration files for ${name}..."
	migrate create -seq -ext=.sql -dir ./migrations ${name}

.PHONY: migration/up
migration/up:
	migrate -path ./migrations -database ${DB_DSN} up