include .env
export $(shell sed 's/=.*//' .env)

http:
	go run cmd/hopper-go/main.go

migrate:
	goose -dir db/migrations up

migrate-down:
	goose -dir db/migrations down
