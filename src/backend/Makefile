SERVICE_NAME = wintrigue
MIGRATION = trigger
TECH = "cmd/tech/main.go"
MIGRATIONS_DIR = "deployments/migrations/db_migrations"
DSN = postgres://$(SERVICE_NAME):password@127.0.0.1:5435/$(SERVICE_NAME)?sslmode=disable

install-goimports:
	go install golang.org/x/tools/cmd/goimports@latest

goimports:
	goimports -w .

gen-server: _gen-server goimports

_gen-server:
	echo ${PWD}
	docker run --rm \
      -v ${PWD}:/backend openapitools/openapi-generator-cli generate \
      -i /backend/api/swagger.yaml \
      -g go-server \
      -o /backend/internal/server/ \
      --additional-properties=outputAsLibrary=true,sourceFolder=.

tech:
	mkdir -p bin
	go build -o bin/tech $(TECH)

create-migration:
	goose --dir=$(MIGRATIONS_DIR) create $(MIGRATION) sql

migrate-up:
	goose --dir=$(MIGRATIONS_DIR) postgres "$(DSN)" up

migrate-down:
	goose --dir=$(MIGRATIONS_DIR) postgres "$(DSN)" down

run-db:
	sudo -S docker-compose --project-directory deployments up -d

stop-db:
	sudo -S docker-compose --project-directory deployments down
