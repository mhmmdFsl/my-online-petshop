.PHONY: clean test security build run

APP_NAME = pet-product
BUILD_DIR = $(PWD)/build
MIGRATIONS_FOLDER = $(PWD)/db/migrations/
DATABASE_URL = postgres://postgres:secret@localhost:5433/postgres?sslmode=disable

clean:
	rm -rf ./build

security:
	gosec -quiet ./...

test: security
	go test -v -timeout 30s -coverprofile=cover.out -cover ./...
	go tool cover -func=cover.out

build: clean test
	CGO_ENABLED=0 go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) main.go

run: swag build
	$(BUILD_DIR)/$(APP_NAME)

docker.run: docker.network swag docker.fiber

docker.network:
	docker network inspect dev-network >/dev/null 2>&1 || \
	docker network create -d bridge dev-network

docker.fiber.build:
	docker build -t fiber .

docker.fiber: docker.fiber.build
	docker run --rm -d \
		--name dev-fiber \
		--network dev-network \
		-p 5000:5000 \
		fiber

docker.fiber.host: docker.fiber.build
	docker run --rm -d \
		--name dev-fiber \
		--network host \
		fiber
		
docker.mongo:
	docker run --rm -d \
		--name dev-mongo \
		--network dev-network \
		-p 27017:27017 \
		mongo

docker.stop: docker.stop.fiber docker.stop.postgres

docker.stop.fiber:
	docker stop dev-fiber

swag:
	swag init