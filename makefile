DB_DSN := "postgres://postgres:123123@localhost:5431/postgres?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN)

migrate-new: #создание новых миграций
	migrate create -ext sql -dir ./migrations ${NAME}

migrate-up: #применение миграций
	$(MIGRATE) up

migrate-down: #откат миграций
	$(MIGRATE) down

migrate-v: #версия
	$(MIGRATE) version

migrate-f: #форс
	$(MIGRATE) force $V

run: #запуск приложения
	go run cmd/app/main.go

gen-tasks: #кодогенерация
	oapi-codegen -config openapi/.openapi -include-tags tasks -package tasks openapi/openapi.yaml > ./internal/web/tasks/api.gen.go

gen-users: #кодогенерация
	oapi-codegen -config openapi/.openapi -include-tags users -package users openapi/openapi.yaml > ./internal/web/users/api.gen.go

lint: #линтер
	golangci-lint run --out-format=colored-line-number
