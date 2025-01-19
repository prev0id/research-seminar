DB_CONTAINER_NAME = calendar_app_postgres
DB_IMAGE = postgres:17.2
DB_PORT = 5432
DB_VOLUME = calendar_app_pgdata
DB_USER = postgres
DB_PASSWORD = password
DB_NAME = calendar_app
APP_ADDR = :8080
DB_CONN = postgres://$(DB_USER):$(DB_PASSWORD)@localhost:$(DB_PORT)/$(DB_NAME)?sslmode=disable
MIGRATIONS_DIR = migrations

dev:
	APP_ADDR=$(APP_ADDR) DB_CONN=$(DB_CONN) \
	go run ./cmd/main.go

start-postgres:
	docker run --rm -d \
		--name $(DB_CONTAINER_NAME) \
		-p $(DB_PORT):5432 \
		-e POSTGRES_USER=$(DB_USER) \
		-e POSTGRES_PASSWORD=$(DB_PASSWORD) \
		-e POSTGRES_DB=$(DB_NAME) \
		-v $(DB_VOLUME):/var/lib/postgresql/data \
		$(DB_IMAGE)

remove-postgres:
	docker rm -f $(DB_CONTAINER_NAME)
	docker volume rm -f $(DB_VOLUME)

migrate:
		goose -dir $(MIGRATIONS_DIR) postgres "$(DB_CONN)" up
