# Build the application
all: build

build:
	@echo "Building..."

	@go build -o main cmd/core/main.go

# Run the application
run:
	@go run cmd/core/main.go

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main

# Live reload
watch:
	@if command -v air > /dev/null; then \
		RUN_MIGRATIONS=1 air; \
		echo "Watching...";\
	else \
		echo "Air is not installed. Run 'go install github.com/air-verse/air@latest' to install Air";\
	fi

sqlc-generate:
	@sqlc generate

templ-generate:
	@templ generate

create-migration:
	@read -p "Give migration a name: " migrationname; \
	migrate create -ext sql -dir internal/db/migrations -seq $$migrationname; \
	echo "Created migration with name $$migrationname. Update sql files."

.PHONY: all build run clean watch sqlc-generate templ-generate create-migration
