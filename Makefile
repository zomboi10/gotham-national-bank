postgres:
	docker run --name postgres17 -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:17-alpine

postgresstop:
	docker stop postgres17

createdb:
	docker exec -it postgres17 createdb --username=root --owner=root gotham_bank

dropdb:
	docker exec -it postgres17 dropdb gotham_bank

migrateup:
	migrate -path ./migrations -database "postgresql://root:secret@localhost:5433/gotham_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path ./migrations -database "postgresql://root:secret@localhost:5433/gotham_bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path ./migrations -database "postgresql://root:secret@localhost:5433/gotham_bank?sslmode=disable" -verbose down 1

server:
	go run main.go

.PHONY: postgres postgresstop createdb dropdb migrateup migratedown migratedown1 server