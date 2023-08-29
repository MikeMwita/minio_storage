.PHONY: clean critic security lint test build run

APP_NAME = filtronicEDAMS
BUILD_DIR = $(PWD)/build
MIGRATIONS_FOLDER = $(PWD)/platform/migrations
DATABASE_URL = postgres://postgres:password@cgapp-postgres/postgres?sslmode=disable

clean:
	rm -rf ./build

critic:
	gocritic check -enableAll ./...

security:
	gosec ./...

lint:
	golangci-lint run ./...

test: clean critic security lint
	go test -v -timeout 30s -coverprofile=cover.out -cover ./...
	go tool cover -func=cover.out

postgres:
	docker run --name filtronicdb -p 5432:5432 -e POSTGRES_USER=filtronic -e POSTGRES_PASSWORD=secret -d postgres:16beta3

createdb:
	docker exec -it filtronicdb createdb --username=filtronic --owner=filtronic edms
#
#docker-compose:
#	docker-compose -f docker_compose.yml up

docker.run: docker.network   docker.fiber

docker.network:
	docker network inspect dev-network >/dev/null 2>&1 || \
	docker network create -d bridge dev-network

docker.fiber.build:
	docker build -t fiber .

docker.fiber: docker.fiber.build
	docker run --rm -d \
		--name cgapp-fiber \
		--network dev-network \
		-p 5000:5000 \
		fiber


docker.stop: docker.stop.fiber

docker.stop.fiber:
	docker stop cgapp-fiber






.PHONY: docker.run docker.network docker.fiber.build docker.fiber docker.stop docker.stop.fiber




