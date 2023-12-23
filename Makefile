LOCAL_MIGRATION_DIR = ./migrations
LOCAL_MIGRATION_DSN = "host=localhost port=5432 dbname=metrics-collection-service user=postgres password=postgres sslmode=disable"

.PHONY: local-migration-up
local-migration-up:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} up -v

.PHONY: local-migration-down
local-migration-down:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} down -v

.PHONY: local-migration-status
local-migration-status:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} status -v

.PHONY: docker-compose-up
docker-compose-up:
	docker-compose up -d

.PHONY: run
run:
	CONFIG_PATH="config/config.yaml" PG_DSN='postgresql://postgres:postgres@localhost:5432/metrics-collection-service' go run cmd/main.go


.PHONY: test
test:
	go test ./...

.PHONY: test-v-cover
test-v:
	go test -v ./...

.PHONY: test-coverage
test-coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html


.PHONY: lint
lint:
	golangci-lint run