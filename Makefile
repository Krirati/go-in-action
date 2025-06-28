g_migration_%:
	goose -dir=db/migrate create $* sql

db_migrate:
	goose -dir=db/migrate -allow-missing postgres "host=localhost port=5432 user=postgres password=postgres dbname=learning sslmode=disable" up

sqlc:
	sqlc generate

server:
	go run main.go

PHONY: g_migration db_migrate sqlc