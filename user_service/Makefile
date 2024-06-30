CURRENT_DIR=$(shell pwd)

APP=$(shell basename "${CURRENT_DIR}")
APP_CMD_DIR=${CURRENT_DIR}/cmd

TAG=latest
ENV_TAG=latest

gen-proto-module:
	./scripts/gen_proto.sh "${CURRENT_DIR}"

migration-up:
	migrate -path ./migrations/postgres -database 'postgres://ilyosbektemirov:12345@0.0.0.0:5432/ms_go_content_service?sslmode=disable' up

migration-down:
	migrate -path ./migrations/postgres -database 'postgres://ilyosbektemirov:12345@0.0.0.0:5432/ms_go_content_service?sslmode=disable' down

build:
	CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -o "${CURRENT_DIR}/bin/${APP}" "${APP_CMD_DIR}/main.go"

run:
	go run cmd/main.go
