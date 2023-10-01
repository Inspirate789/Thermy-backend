cli := $(go env GOBIN)/thermy-cli
CLI_SRC := $(wildcard cmd/thermy-cli/*.go)

.PHONY: run swagger cli test setup_redis

run: swagger
	./build/deploy.sh

swagger:
	swag fmt
	swag init --parseDependency --parseInternal --parseDepth 1 -g cmd/backend/main.go -o swagger/

cli: $(CLI_SRC)
	go install ./cmd/thermy-cli

test:
	godotenv -f ./.env go test ./...

setup_redis:
	godotenv -f ./.env go run ./db/redis
