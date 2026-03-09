# Target execution
.PHONY: run build migrate seed migrate-create test tidy clean docker-up docker-down help

# Default target
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  run          Run the API server"
	@echo "  build        Build the API server binary"
	@echo "  migrate-up   Run database migrations"
	@echo "  migrate-down Run database migrations"
	@echo "  seed         Run database seeders"
	@echo "  migrate-create Create a new migration file (usage: make migrate-create name=migration_name)"
	@echo "  test         Run all tests"
	@echo "  tidy         Run go mod tidy"
	@echo "  clean        Clean build artifacts"
	@echo "  docker-up    Start services with docker-compose"
	@echo "  docker-down  Stop services with docker-compose"

run:
	go run cmd/api/main.go

build:
	mkdir -p bin
	go build -o bin/api cmd/api/main.go

migrate-up:
	go run cmd/migrate/main.go up

migrate-down:
	go run cmd/migrate/main.go down

seed:
	go run cmd/seeder/main.go

migrate-create:
	go run cmd/migrate/main.go create $(name)

tidy:
	go mod tidy

clean:
	rm -rf bin/
	rm -rf tmp/

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down
