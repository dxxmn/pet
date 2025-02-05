DB_DSN := "postgres://postgres:123123@localhost:5431/postgres?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN)

migrate-new:
	migrate create -ext sql -dir ./migrations ${NAME}

migrate:
	$(MIGRATE) up

migrate-v:
	$(MIGRATE) version

migrate-f:
	$(MIGRATE) force $V

migrate-down:
	$(MIGRATE) down

run:
	go run cmd/app/main.go