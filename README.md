# go-crudapi\
# Tototal create container manual
docker container create --name database -p 5433:5432 -e POSTGRES_PASSWORD=123456  postgres:alpine3.14

# Migrate table
migrate -path db/migrations -database "postgresql://root:123456@localhost:5433/dbcrud?sslmode=disable" -verbose up