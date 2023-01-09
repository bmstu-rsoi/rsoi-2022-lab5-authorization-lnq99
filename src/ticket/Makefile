DB_DRIVER=postgresql
USERNAME=program
PASSWORD=test
HOST=localhost
PORT=5432
DB_NAME=tickets

DB_URL="$(DB_DRIVER)://$(USERNAME):$(PASSWORD)@$(HOST):$(PORT)/$(DB_NAME)?sslmode=disable"

migrate_create:
	migrate create -ext sql -dir db -seq init
	migrate create -ext sql -dir db -seq seed

migrate_up:
	migrate -path db -database $(DB_URL) -verbose up

migrate_down:
	migrate -path db -database $(DB_URL) -verbose down

sqlc:
	sqlc generate
