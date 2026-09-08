# Avito.Кухня

MVP backend-сервиса доставки еды: рестораны регистрируются, выкладывают меню
и принимают заказы пользователей через единый API платформы.

- Клиентский API — для web-клиента (`X-User-Id`).
- Партнёрский API — для заведений (`X-Api-Key`).

> Аутентификация пользователей и сам клиент по условиям задания **не реализуются**.

---

## Содержание

- [Avito.Кухня](#avitoкухня)
  - [Содержание](#содержание)
  - [Стек](#стек)
  - [Структура репозитория](#структура-репозитория)
  - [Быстрый старт](#быстрый-старт)
  - [Архитектура (C4 L2)](#архитектура-c4-l2)
  - [API](#api)
  - [Контракт ресторана (callback)](#контракт-ресторана-callback)
    - [Запрос от Avito → ресторан](#запрос-от-avito--ресторан)
    - [Ответ ресторана → Avito](#ответ-ресторана--avito)
    - [Требования к реализации эндпоинта](#требования-к-реализации-эндпоинта)
    - [Что делать с заказами после подтверждения](#что-делать-с-заказами-после-подтверждения)
  - [Схема БД](#схема-бд)
  - [Жизненный цикл заказа](#жизненный-цикл-заказа)
    - [Создание](#создание)
    - [Обновление статуса](#обновление-статуса)
  - [CJM](#cjm)
    - [Сценарий пользователя](#сценарий-пользователя)
    - [Сценарий ресторана](#сценарий-ресторана)
  - [Упрощения MVP](#упрощения-mvp)
  - [Возможные улучшения:](#возможные-улучшения)
  - [Линтеры](#линтеры)

---

## Стек

- Go 1.25: echo, sqlx, pgx, golang-migrate
- PostgreSQL 14

## Структура репозитория

- `avito/` — основной сервис (HTTP API + миграции)
- `deliciousfood/` — пример ресторана-партнёра
- `config/.env.dev` — переменные окружения для dev
- `docker-compose.dev.yml` — оркестрация
- `docs/` — CJM-диаграммы (PUML + PNG)
- `Makefile` — `up` / `down` / `logs` / `ps` / `build` / `clean`
- OpenAPI: [`avito/api/openapi.yaml`](./avito/api/openapi.yaml)
- Линтер: [`avito/.golangci.yml`](./avito/.golangci.yml)

## Быстрый старт

```bash
make up        # postgres + миграции + avito + deliciousfood
make logs      # логи всех контейнеров
make down      # остановить (с сохранением данных)
make clean     # остановить и удалить volumes

# Или
docker compose --env-file config/.env.dev -f docker-compose.dev.yml up -d
```

После старта:
- **Avito.Кухня** — `http://localhost:8080`
- **deliciousfood** — `http://localhost:8081`
- **Postgres** — `localhost:5432` (`postgres/postgres`, db `avito`)

Получить `api-key` deliciousfood можно через логи контейнера
```bash
make logs
# Или
docker compose --env-file config/.env.dev -f docker-compose.dev.yml logs deliciousfood
```
Для тестирования рекомендую использовать [swagger editor](https://forge.etsi.org/swagger/editor/). OpenAPI: [`avito/api/openapi.yaml`](./avito/api/openapi.yaml).

## Архитектура (C4 L2)

![C4 L2 — контейнеры](./docs/c4-l2.png)

Исходник: [`docs/architecture.puml`](./docs/architecture.puml).

**Слои внутри `avito`:** `transport` (Echo) → `service` (бизнес-логика) →
`repository` (sqlx) → Postgres. Внутри также `pkg/restaurantapi` — HTTP-клиент
для синхронных callback'ов к ресторанам с ретраями и таймаутом.

## API

Полная спецификация — [`avito/api/openapi.yaml`](./avito/api/openapi.yaml).

| Метод & путь | Кто вызывает | Auth |
| --- | --- | --- |
| `GET /restaurants` | клиент | — |
| `POST /restaurants` | ресторан (регистрация) | — |
| `GET /restaurants/{id}` | клиент | — |
| `PATCH /restaurants/{id}` | ресторан | `X-Api-Key` |
| `GET /restaurants/{id}/menu` | клиент | — |
| `POST /menu` | ресторан | `X-Api-Key` |
| `PATCH /menu/{id}` | ресторан | `X-Api-Key` |
| `POST /orders` | клиент | `X-User-Id` |
| `GET /orders` | клиент | `X-User-Id` |
| `GET /orders/{id}` | клиент | `X-User-Id` |
| `GET /orders/{id}/history` | клиент | `X-User-Id` |
| `PATCH /orders/{id}/status` | ресторан | `X-Api-Key` |
| `GET /health` | мониторинг | — |

Все ошибки возвращаются в формате `{ "status": "error", "message": "..." }`,
коды: `400 / 401 / 403 / 404 / 409 / 500`.

## Контракт ресторана (callback)

Когда пользователь оформляет заказ, Avito.Кухня **синхронно** вызывает
`callback_url`, который ресторан указал при регистрации. Ресторан **сам
реализует** этот эндпоинт и решает, принять заказ или отклонить.
### Запрос от Avito → ресторан

`POST {callback_url}` (например, `http://deliciousfood:8081/orders`), `Content-Type: application/json`:

```json
{
  "order": {
    "id": 1,
    "user_id": 42,
    "restaurant_id": 1,
    "status": "pending",
    "delivery_address": "г. Казань, ул. Баумана, 20",
    "comment": "Позвоните за 10 минут",
    "total_price": "1240.00",
    "created_at": "2026-09-07T10:00:00Z",
    "updated_at": "2026-09-07T10:00:00Z"
  },
  "items": [
    {
      "id": 1,
      "order_id": 1,
      "menu_item_id": 3,
      "quantity": 2,
      "price": "590.00"
    }
  ]
}
```

### Ответ ресторана → Avito

| Код | Когда возвращать | Что делает Avito |
| --- | --- | --- |
| **201 Created** + `{"status":"success","message":"…"}` | Заказ принят | Ставит заказу статус `confirmed`, пишет `message` в `orders.restaurant_message` |
| **409 Conflict** + `{"status":"error","message":"причина"}` | Заказ отклонён | Ставит заказу статус `rejected`, текст причины возвращается пользователю в `409` |
| любой другой код / таймаут | Техническая проблема | Avito интерпретирует как отказ (см. ниже) |

### Требования к реализации эндпоинта

1. **Идемпотентность.** Avito делает до 3 ретраев с экспоненциальным backoff
   (1 c → 2 c → 4 c) и общим таймаутом 2 c на попытку. Ресторан должен
   корректно отвечать одинаково на повторные вызовы с тем же `order.id`.
2. **Быстрый ответ.** `deliciousfood` отвечает сразу `201`/`409`, а
   дальнейшую обработку (смену статусов) делает **сам** через
   `PATCH /orders/{id}/status`. Не блокируйте HTTP-вызов Avito тяжёлой
   логикой.
3. **Точка интеграции — `callback_url` целиком.** Это может быть `http://…/orders`
   или `https://…/webhooks/avito` — что угодно, лишь бы POST с JSON работал
   и помещался в 2 c.
4. **Безопасность.** В MVP callback не аутентифицируется (Avito не передаёт
   подпись). В проде стоит добавить HMAC-подпись в заголовке и проверку на
   стороне ресторана.
5. **Сетевая досягаемость.** `callback_url` должен быть доступен из
   Docker-сети `avito-network`. В `docker-compose.dev.yml` ресторан
   присутствует в той же сети и обращается к Avito по имени сервиса `avito`,
   а Avito вызывает его по `http://deliciousfood:8081/orders`.

### Что делать с заказами после подтверждения

После `201` ресторан сам отвечает за жизненный цикл заказа и пушит статусы
через `PATCH /orders/{id}/status` с заголовком `X-Api-Key`:

```
cooking → ready → delivering → delivered
```

В любой момент можно отклонить, передав `rejected` (Avito не даст изменить
заказ в финальных статусах `rejected` / `delivered`).

## Схема БД

![Схема БД](./docs/schema.png)

Исходник миграций — [`avito/migrations/`](./avito/migrations) (5 файлов
`*.up.sql` / `*.down.sql`).

Ключевые решения:

- Цена в `order_items` — **снапшот** цены из `menu_items` на момент заказа.
- История статусов (`order_status_history`) ведётся **триггером** PostgreSQL
  `update_order_status_history_trigger` — любой INSERT/UPDATE в `orders`
  автоматически порождает запись в истории (аудит без логики в коде).

## Жизненный цикл заказа

### Создание

![Создание заказа](./docs/order_create.png)

### Обновление статуса

![Обновление статуса](./docs/order_update.png)

| Статус | Кто выставляет | Через что |
| --- | --- | --- |
| `pending` | Avito | автоматически при `POST /orders` |
| `confirmed` / `rejected` | Avito | по ответу ресторана на callback |
| `cooking` / `ready` / `delivering` / `delivered` | ресторан | `PATCH /orders/{id}/status` |

Запрещено менять `pending` (его ставит только API) и заказы в финальных
статусах (`rejected`, `delivered`). Валидация порядка переходов между
нефинальными статусами в MVP не выполняется — ресторан считается доверенным.

## CJM

Исходники: [`docs/user.puml`](./docs/user.puml),
[`docs/restaurant.puml`](./docs/restaurant.puml).

Текстовые описания — [`docs/user_flow.txt`](./docs/user_flow.txt),
[`docs/restaurant_flow.txt`](./docs/restaurant_flow.txt).

### Сценарий пользователя
![Сценарий пользователя](./docs/cjm_user.png)

### Сценарий ресторана
![Сценарий ресторана](./docs/cjm_restaurant.png)

## Упрощения MVP

- Возможная потеря данных:    
  Пользователь создает заказ -> заказ сохраняется в БД(status `pending`) -> ресторан отклоняет заказ -> приложение падает/бд недоступна -> "мертвый" заказ со статусом `pending`.  
  Возможное решение:  
  Фоновый процесс для отслеживания "мертвых" заказов.
- `X-User-Id` вместо полноценной auth пользователей.
- API-ключ ресторана хранится в открытом виде в БД.
- Нет rate-limit на callback к ресторану; есть только
  3 ретрая с экспоненциальным backoff и таймаутом 2 c.
- `deliciousfood` рандомно отклоняет 50% заказов — чтобы в демо были видны
  оба сценария CJM.
- Нет валидации переходов между нефинальными статусами.
- Нет пагинации в списках.
- Логи в stdout через `log.Printf`.
- Нет юнит-/интеграционных тестов.

## Возможные улучшения: 
- JWT вместо `X-User-Id`
- Хэширование API-ключа ресторана
- Брокер сообщений для событий `order.status_changed` 
- Prometheus + Grafana + Loki для мониторинга.

## Линтеры

`avito/.golangci.yml` включает `errcheck`, `govet`, `staticcheck`, `unused`,
`bodyclose`, `contextcheck`, `errorlint`, `errname`, `noctx`, `sqlclosecheck`,
`unconvert`, `unparam`, `misspell`, `gocritic`, `gosec`, `prealloc`,
`copyloopvar`, `dupl`, `goconst`. Форматтеры: `gofmt` + `goimports`.

```bash
cd avito && golangci-lint run ./...
cd deliciousfood && golangci-lint run ./...
```
