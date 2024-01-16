.PHONY: clean critic security lint test build run

APP_NAME = EDAMS
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


docker-compose:
	docker-compose up --build
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



proto:
	protoc --proto_path=gapi/protos/mutation \
      --go_out=gapi/pb/mutation_gen \
      --go_opt=paths=source_relative \
      --go-grpc_out=gapi/pb/mutation_gen \
      --go-grpc_opt=paths=source_relative \
      --grpc-gateway_out=gapi/pb/mutation_gen \
      --grpc-gateway_opt=paths=source_relative \
      gapi/protos/mutation/*.proto


evans:
	evans --host localhost --port 9090 -r repl


.PHONY: docker.run docker.network docker.fiber.build docker.fiber docker.stop docker.stop.fiber




