# Load variables from .env file (if it exists)
ifneq (,$(wildcard .env))
    include .env
    export $(shell sed 's/=.*//' .env)
endif

MIGRATIONS_DIR := $(shell pwd)/migrations

create-db:
	docker exec -it authdb creatdb auth

migrate-create:
	@if [ -z "$(name)" ]; then \
		echo "❌ Please specify a migration name: make migrate-create name=user"; \
		exit 1; \
	fi
	migrate create -ext sql -seq -digits 4 $(name) -dir $(MIGRATIONS_DIR)

migrate-up:
	migrate -database $(DB_ADDRESS) -source $(MIGRATIONS_DIR) -verbose up

migrate-down:
	migrate -database $(DB_ADDRESS) -source $(MIGRATIONS_DIR) -verbose down
