include .env

VERSION = 1.0.0

migration_up: 
	migrate -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:5432/postgres?sslmode=disable" -path internal/repositories/migrations up

migration_down: 
	migrate -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:5432/postgres?sslmode=disable" -path internal/repositories/migrations down

test:
	go test -race -v ./...

version:
	echo ${VERSION}

lint:
	golangci-lint run --config .linterconf.yml