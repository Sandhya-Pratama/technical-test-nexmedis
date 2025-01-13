# Load environment variables
ifneq (,$(wildcard .env))
    include .env
    export $(shell sed 's/=.*//' .env)
endif

# PostgreSQL container
postgres:
    docker run --name postgres --network bank-network \
    -e POSTGRES_USER=$(POSTGRES_USER) \
    -e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) \
    -p $(POSTGRES_PORT):5432 -d postgres:$(POSTGRES_VERSION)

# Create database
createdb:
    docker exec -it postgres createdb --username=$(POSTGRES_USER) --owner=$(POSTGRES_USER) $(POSTGRES_DB)

# Drop database
dropdb:
    docker exec -it postgres dropdb $(POSTGRES_DB)

# Create migrations
createmigrations:
    goose create create_users_table sql
    goose create create_orders_table sql

# Migrate up
migrateup:
    goose -dir db/migrations postgres \
    "user=$(POSTGRES_USER) password=$(POSTGRES_PASSWORD) dbname=$(POSTGRES_DB) sslmode=disable" up

# Migrate down
migratedown:
    goose -dir db/migrations postgres \
    "user=$(POSTGRES_USER) password=$(POSTGRES_PASSWORD) dbname=$(POSTGRES_DB) sslmode=disable" down

# Run server
server:
    go run cmd/main.go

.PHONY: postgres createdb dropdb createmigrations migrateup migratedown server
