include app.env
postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=${POSTGRES_USER} -e POSTGRES_PASSWORD=${POSTGRES_PASSWORD} -d postgres:12-alpine

redis:
	docker run --name redis3 -p 6379:6379 -d redis:6.2-alpine3.13 redis-server --requirepass ${REDIS_PASS}

createdb:
	docker exec -it postgres12 createdb --username=${POSTGRES_USER} --owner=${POSTGRES_USER} ${DB_NAME}

dropdb:
	docker exec -it postgres12 dropdb -U ${POSTGRES_USER} ${DB_NAME}

migrate_up:
	go run main.go migrate_up

ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

sqlc_generate:
	docker run --rm -v $(ROOT_DIR):/src -w /src kjconroy/sqlc generate

generate_docs:
	go run main.go generate_docs

test:
	go run main.go migrate_up -t| true #ignore exit 0
	go test ./...  -v -coverprofile cover.out
	@echo "================================================"
	@echo "Coverage"
	go tool cover -func cover.out
	@echo "Cleaning..."
	go run main.go dropdb -t
# @echo "docker run --rm -v ${d}:/src -w /src kjconroy/sqlc generate"
# createtestdb:
# 	docker exec -it postgres12 createdb --username=sponge --owner=sponge twtTest

# droptestdb:
# 	docker exec -it postgres12 dropdb -U sponge twtTest

.PHONY: postgres createdb dropdb migrate_up sqlc_generate test redis # createtestdb droptestdb
