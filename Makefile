GOOSE_DIR=sql/schemas
GOOSE_DRIVER?=postgres
DB_NAME=ChatApp
DB_USER=root 
DB_SSLMODE=disable
PG_HOST = DB_SOURCE="postgresql://root:secret@postgres17:5432/ChatApp?sslmode=disable"

postgres:
	docker run --name postgres17 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:17rc1-alpine3.20

startdk:
	docker start postgres17

createdb:
	docker exec -it postgres17 createdb --username=root --owner=root ChatApp

upgo:
	goose -dir $(GOOSE_DIR) postgres postgres://root:secret@localhost:5432/ChatApp?sslmode=disable up

builddk:
	docker build . -t chatapp

network-create:
	docker network create chatapp-network
	docker network connect chatapp-network postgres17

rundk:
	docker run --name chatapp --network chatapp-network -p 8050:8050 -e GIN_MODE=release -e $(PG_HOST) chatapp:latest

##########

# You can skip here
server:
	go run cmd/main.go

createtable:
	goose -dir $(GOOSE_DIR) create $(name) sql

downgo:	
	goose -dir $(GOOSE_DIR) postgres postgres://root:secret@localhost:5432/ChatApp?sslmode=disable down

dropdb:
	docker exec -it postgres17 dropdb --username=root ChatApp

sqlc:
	sqlc generate

module:
	touch internal/repo/$(name).repo.go
	touch internal/service/$(name).service.go
	touch internal/controller/$(name).controller.go
	touch internal/router/$(name).router.go
# 

.PHONY: server startdk createdb dropdb createtable upgo downgo sqlc module