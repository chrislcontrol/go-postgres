PROJECT_NAME := go-postgres

# Containers:
# Postgres Local
run-postgres:
	docker start $(PROJECT_NAME)-postgres 2>/dev/null || docker run --name $(PROJECT_NAME)-postgres -p 5432:5432 -e POSTGRES_PASSWORD=postgres -d postgres:15-alpine

containers-down:
	docker stop $$(docker ps -aq)

containers-up: run-postgres

containers-reset: containers-down containers-up


run:
	go run cmd/main.go
