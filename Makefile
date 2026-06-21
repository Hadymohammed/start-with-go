.PHONY: generate-sqlc migrate-up migrate-down swag run build

generate-sqlc:
	go run github.com/sqlc-dev/sqlc/cmd/sqlc generate

migrate-up:
	go run github.com/pressly/goose/v3/cmd/goose -dir sql/migrations postgres "$(DATABASE_URL)" up

migrate-down:
	go run github.com/pressly/goose/v3/cmd/goose -dir sql/migrations postgres "$(DATABASE_URL)" down

swag:
	swag init -g cmd/api/main.go -o docs

run:
	go run ./cmd/api

build:
	go build -o bin/api ./cmd/api
