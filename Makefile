GOOSE_DIR=sql/schemas
GOOSE_DRIVER?=postgres
DB_NAME=ChatApp
DB_USER=root 
DB_SSLMODE=disable

postgres:
	docker run --name postgres17 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=yourchoice -d postgres:17rc1-alpine3.20

server:
	go run cmd/main.go

startdk:
	docker start postgres17

createdb:
	docker exec -it postgres17 createdb --username=root --owner=root ChatApp

createtable:
	goose -dir $(GOOSE_DIR) create $(name) sql

upgo:
	goose -dir $(GOOSE_DIR) postgres postgres://root:secret@localhost:5432/ChatApp?sslmode=disable up

# You can skip here
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