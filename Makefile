
migrateup:
	migrate -path="migrations"  -database="postgres://saldop_api:saldop_api_password@0.0.0.0:5432/saldop?sslmode=disable" up

migratedown:
	migrate -path="migrations"  -database="postgres://saldop_api:saldop_api_password@0.0.0.0:5432/saldop?sslmode=disable" down

database:
	docker compose up database -d

nukedatabase:
	docker compose down database
	docker volume rm saldop_database_data