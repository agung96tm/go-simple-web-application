include .envrc

.PHONY: proto_build
proto_build:
	protoc --go_out=./internal/proto --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=./internal/proto ./internal/proto/*.proto


.PHONY: run/rpc
run/rpc:
	go run ./cmd/rpc -db-dsn=${DB_DSN}

.PHONY: run/rest
run/rest:
	go run ./cmd/rest


.PHONY: migration/new
migration/new:
	@echo "Creating migration files for ${name}..."
	migrate create -seq -ext=.sql -dir ./migrations ${name}

.PHONY: migration/up
migration/up:
	migrate -path ./migrations -database ${DB_DSN} up