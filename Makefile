include .env

run:
	go run cmd/main.go

tidy:
	go mod tidy

gsql:
	sqlc generate

create-migration:
	migrate create -ext sql -dir migrations "${name}"

migrate-up:
	migrate -path migrations -database "postgresql://${DATABASE_USER}:${DATABASE_PASSWORD}@${DATABASE_HOST}:${DATABASE_PORT}/${DATABASE_DB}?sslmode=disable" up

migrate-down:
	migrate -path migrations -database "postgresql://${DATABASE_USER}:${DATABASE_PASSWORD}@${DATABASE_HOST}:${DATABASE_PORT}/${DATABASE_DB}?sslmode=disable" down

