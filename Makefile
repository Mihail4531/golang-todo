include .env
export

export PROJECT_ROOT= $(shell pwd)

env-up:
	docker-compose up -d todoapp-postgres

env-down:	
	docker-compose down todoapp-postgres 
env-port-forward:
	@docker-compose up -d port-forwarder
env-port-close:
	@docker-compose down  port-forwarder
env-cleanup:
	@echo "Удаление данных базы данных..."
	docker compose down
	@rm -rf out/pgdata
	@mkdir -p out/pgdata
	@echo "Данные очищены."

migrate-create:
	@if [ -z "$(seq)" ]; then \
		echo "Ошибка: переменная seq не задана! Используй: make migrate-create seq=имя_миграции"; \
		exit 1; \
	fi;
	docker compose run --rm \
		-e GOOSE_COMMAND=create \
		-e GOOSE_COMMAND_ARG="$(seq) sql" \
		-e GOOSE_EXTRA_ARGS="-s" \
		todoapp-postgres-migrate

migrate-up:
	docker compose run --rm \
		-e GOOSE_COMMAND=up \
		todoapp-postgres-migrate

migrate-down:
	docker compose run --rm \
		-e GOOSE_COMMAND=down \
		todoapp-postgres-migrate