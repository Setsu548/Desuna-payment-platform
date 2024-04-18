postgres:
	docker compose up -d
	# docker pull postgres:latest
	# docker rm -f postgres16
	# docker run --name postgres16 --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=Control123 -d postgres:latest

createdb:
	docker exec -it postgres-db createdb --username=root --owner=root bank_simulator

dropdb:
	docker exec -it postgres16 dropdb bank_simulator

migrateup:
	migrate -path db/migration -database "postgresql://root:Control123@localhost:5432/bank_simulator?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:Control123@localhost:5432/bank_simulator?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:Control123@localhost:5432/bank_simulator?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:Control123@localhost:5432/bank_simulator?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb  -destination db/mock/store.go github.com/Petatron/bank-simulator-backend/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc test server mock
# run-in-sequence: postgres dropdb createdb migratedown migrateup

