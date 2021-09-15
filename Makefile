postgres:
	docker run --name database -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=123456  postgres:alpine3.14

createdb:
	docker exec -it database createdb --username=root --owner=root dbcrud

dropdb:
	docker exec -it database dropdb dbcrud

migrateup:
	migrate -path db/migrations -database "postgresql://root:123456@localhost:5433/dbcrud?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://root:123456@localhost:5433/dbcrud?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb migrateup 

