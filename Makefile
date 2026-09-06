.PHONY: help up down restart logs ps build clean

help:
	@echo "Avito.Кухня"
	@echo ""
	@echo "  make up        - запустить все сервисы"
	@echo "  make down      - остановить все сервисы"
	@echo "  make restart   - перезапустить все сервисы"
	@echo "  make logs      - показать логи всех сервисов"
	@echo "  make ps        - показать статус сервисов"
	@echo "  make build     - пересобрать образы"
	@echo "  make clean     - остановить и удалить volumes"

up:
	docker compose --env-file config/.env.dev -f docker-compose.dev.yml up -d

down:
	docker compose --env-file config/.env.dev -f docker-compose.dev.yml down

restart:
	docker compose --env-file config/.env.dev -f docker-compose.dev.yml restart

logs:
	docker compose --env-file config/.env.dev -f docker-compose.dev.yml logs -f

ps:
	docker compose --env-file config/.env.dev -f docker-compose.dev.yml ps

build:
	docker compose --env-file config/.env.dev -f docker-compose.dev.yml build

clean:
	docker compose --env-file config/.env.dev -f docker-compose.dev.yml down -v
