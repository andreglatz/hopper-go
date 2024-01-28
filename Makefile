include .env
export $(shell sed 's/=.*//' .env)

http:
	go run cmd/hopper-go/main.go

create-migration:
	goose -dir db/migrations create $(name) sql


migrate:
	goose -dir db/migrations up

migrate-down:
	goose -dir db/migrations down
