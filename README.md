# Product API

Small Go service for managing products with PostgreSQL and Docker Compose.

## Prerequisites

- Docker
- Docker Compose
- Go 1.26+ (only required for local execution, not Docker)

## Required `.env` file

The application uses `github.com/joho/godotenv` to load environment variables from a `.env` file.
Create a `.env` file in the project root with the following values:

```env
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=products_db
DB_PORT=5432
DB_SSLMODE=disable
PORT=8080
```

> The `.env` file is required for local execution. When using Docker Compose, the `api` service already receives database variables from `docker-compose.yaml`, but you can still use `.env` to override or run the app directly.

## Run with Docker Compose

From the project root:

```bash
docker compose up --build
```

This will start:

- PostgreSQL on port `5432`
- API service on port `8080`

### Stop the services

```bash
docker compose down
```

## Run locally without Docker

1. Create the `.env` file in the project root as shown above.
2. Install dependencies and tidy modules:

```bash
go mod tidy
```

3. Run the API:

```bash
go run cmd/api/main.go
```

The service will start on `http://localhost:8080`.

## Application configuration

The server uses the following environment variables:

- `DATABASE_URL` (optional) — full PostgreSQL DSN string
- `DB_HOST` — database host
- `DB_USER` — database username
- `DB_PASSWORD` — database password
- `DB_NAME` — database database name
- `DB_PORT` — database port
- `DB_SSLMODE` — PostgreSQL SSL mode
- `PORT` — HTTP port for the API

If `DATABASE_URL` is set, it takes precedence over the individual DB_* variables.

## API endpoints

The project registers REST routes via Gin. The service is exposed on `http://localhost:8080`.

Example routes:

- `POST /products`
- `GET /products/:id`
- `PUT /products/:id`
- `DELETE /products/:id`
- `GET /products`

## Notes

- The app performs automatic schema migration on startup using GORM.
- If you change dependencies, run `go mod tidy` again.
