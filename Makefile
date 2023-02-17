BINARY_NAME=bin/app
COMPOSE="docker-compose.yml"

build:
	GOARCH=amd64 go build -o ${BINARY_NAME} main.go

run:
	./${BINARY_NAME}

build_and_run: build run

clean:
	go clean
	rm ${BINARY_NAME}

compose_build:
	docker-compose -f ${COMPOSE} build

compose_run:
	docker-compose -f ${COMPOSE} up

compose_build_and_run: compose_build compose_run

compose_clean:
	docker-compose -f ${COMPOSE} down
	docker volume ls -q | grep highload_architect | xargs docker volume rm

test_api:
	pytest test/api

generate:
	openapi-generator generate \
		-i ./api/openapi.json \
		-g go-server \
		-o ./generated/go_server
	openapi-generator generate \
		-i ./api/openapi.json \
		-g python-prior \
		-o ./generated/python_client
	python3 generated/patch_go_server.py | gofmt | tee "generated/go_server/go/authorize_routes.go"
