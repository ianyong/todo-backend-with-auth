BINARY_PATH=./bin/server/server.out
SERVER_PATH=./cmd/server/main.go
DB_CREATE_PATH=./cmd/database/create/main.go
DB_MIGRATE_PATH=./cmd/database/migrate/main.go
DB_DROP_PATH=./cmd/database/drop/main.go

createdb:
	@echo "Creating database..."
	@go run ${DB_CREATE_PATH}

migratedb:
	@echo "Migrating database..."
	@go run ${DB_MIGRATE_PATH}

dropdb:
	@echo "Dropping database..."
	@go run ${DB_DROP_PATH}

build:
	@echo "Building server at ${BINARY_PATH}..."
	@go build -o ${BINARY_PATH} ${SERVER_PATH}

run:
	@echo "Running server..."
	@go build -o ${BINARY_PATH} ${SERVER_PATH}
	@eval ${BINARY_PATH}

fmt:
	@echo "Formatting source files..."
	@go fmt ./...

lint:
	@echo "Linting source files..."
	@golangci-lint run

lintfix:
	@echo "Linting and automatically fixing source files..."
	@golangci-lint run --fix

clean:
	@echo "Removing object files and cached files..."
	@go clean
	@rm ${BINARY_PATH}

test:
	@echo "Running tests..."
	@GO_ENV=test go test ./...
