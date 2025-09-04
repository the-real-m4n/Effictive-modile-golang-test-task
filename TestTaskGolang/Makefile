DB_URL=postgres://postgres:postgres@127.0.0.1:5433/subscriptions?sslmode=disable

.PHONY: up down migrate-up migrate-down run

up:
	docker compose up -d

down:
	docker compose down -v

migrate-up:
	migrate -path ./migrations -database "$(DB_URL)" up

migrate-down:
	migrate -path ./migrations -database "$(DB_URL)" down

run:
	go run ./cmd/app
