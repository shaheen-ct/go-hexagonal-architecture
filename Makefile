HELP_CMD = grep -E '^[a-zA-Z_-]+:.*?\#\# .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?\#\# "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

init: # all the mandatory generator packaged
	go get github.com/google/wire@latest # DI

env:
	source .env

.PHONY: up-server
up-server: migrate-up ## run server
	go run cmd/server/server.go

.PHONY: up-develop-server
up-develop-server:migrate-up ## run develop server with live loader air
	air

.PHONY: test
test: ## run all tests
	go test ./...

.PHONY: generate
generate: ## run all generators
	go generate ./...

.PHONY: generate-gql
generate-gql: ## regenerate gql resolvers
	go run github.com/99designs/gqlgen generate

tool: ## Add github.com/99designs/gqlgen to your project’s tools.go
	printf '// +build tools\npackage tools\nimport _ "github.com/99designs/gqlgen"' | gofmt > tools.go
	go mod tidy

.PHONY: help
help: ## Show this help
	@${HELP_CMD}

# DB

.PHONY: postgres
postgres:
	docker run --name meetup -e POSTGRES_PASSWORD=123456 -e POSTGRES_USER=root -p 5432:5432 -d postgres:12-alpine

.PHONY: create-db
create-db:
	docker exec -it meetup createdb --username=root --owner=root meetup

.PHONY: migrate-up
migrate-up:
	migrate -path db/pgx/migrations -database "postgres://root:123456@localhost:5432/meetup?sslmode=disable" -verbose up