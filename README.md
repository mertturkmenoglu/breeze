# Breeze

## Prerequisites

- Go 1.23+
- Docker
- Make
- SQLC
- Templ
- Air
- npm

## Setup

- Copy `.env.example` to `.env` and fill in the required variables.
- Run `go mod download` to download the required dependencies.
- Run `make sqlc-generate` to generate the SQL schema.
- Run `make templ-generate` to generate the HTML templates.

## Running

- Run `docker compose up -d` to start the containers.
- Inside `internal/styles` directory, run `npm run watch-css` to watch for changes in the CSS files and automatically compile them.
- Run `templ generate --watch` to watch for changes in Templ files and automatically compile them.
- Run `make watch` to watch for changes in Go files and automatically compile them.
