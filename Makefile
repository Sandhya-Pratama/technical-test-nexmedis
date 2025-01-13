postgres:
	docker run --name postgres --network bank-network -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=sandhya123 -p 5432:5432 -d postgres:16

createdb:
	docker exec -it postgres createdb --username=postgres --owner=postgres e-commerce

dropdb:
	docker exec -it postgres dropdb e-commerce

createmigrations:
	goose -dir db/migrations create create_users_table sql
	goose -dir db/migrations create create_orders_table sql
	goose -dir db/migrations create create_cart_items_table sql
	goose -dir db/migrations create create_order_table sql
	goose -dir db/migrations create create_order_items_table sql

migrateup:
	goose -dir db/migrations postgres "user=postgres password=sandhya123 dbname=e-commerce sslmode=disable" up

migratedown:
	goose -dir db/migrations postgres "user=postgres password=sandhya123 dbname=e-commerce sslmode=disable" down

server:
	go run cmd/main.go

.PHONY: postgres createdb dropdb createmigrations migrateup migratedown server
