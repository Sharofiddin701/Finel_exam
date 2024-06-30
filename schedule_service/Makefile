CURRENT_DIR=$(shell pwd)

gen-proto-module:
	./scripts/gen_proto.sh "${CURRENT_DIR}"

migration-up:
	migrate -path ./migrations/postgres -database 'postgres://user_name:password@0.0.0.0:5432/database?sslmode=disable' up

migration-down:
	migrate -path ./migrations/postgres -database 'postgres://user_name:password@0.0.0.0:5432/database?sslmode=disable' down

run:
	go run cmd/main.go
