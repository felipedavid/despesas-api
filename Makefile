
migrateup:
	migrate -path="migrations"  -database="postgres://saldop_api:saldop_api_password@0.0.0.0:5432/saldop?sslmode=disable" up

migratedown:
	migrate -path="migrations"  -database="postgres://saldop_api:saldop_api_password@0.0.0.0:5432/saldop?sslmode=disable" down

database:
	docker compose up database -d

nukedatabase:
	docker compose down database
	docker volume rm saldop-api_database_data

build:
	go build -o ./bin/api ./cmd

run: build
	./bin/api

generate_translations:
	go generate ./internal/translations/translations.go

test:
	docker build -q --target export-test -o ./out .

image:
	docker build -o out --build-arg VERSION=tururu .
