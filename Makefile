.PHONY: generate-sqlc generate-api migrate-up migrate-down run build

generate-sqlc:
	go run github.com/sqlc-dev/sqlc/cmd/sqlc generate

generate-api:
	go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen \
		--config oapi-codegen.yaml \
		api/tsp-output/schema/openapi.yaml

migrate-up:
	export $$(cat .env | xargs) && go run github.com/pressly/goose/v3/cmd/goose -dir sql/migrations postgres $$DATABASE_URL up

migrate-down:
	export $$(cat .env | xargs) && go run github.com/pressly/goose/v3/cmd/goose -dir sql/migrations postgres $$DATABASE_URL down

run:
	go run ./cmd/api

build:
	go build -o bin/api ./cmd/api
