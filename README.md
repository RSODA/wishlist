# Wishlist

Проект на базе Telegram Web App, который позволяет пользователям
добавлять свои желания и делиться ими с друзьями.

## Требования

- Go 1.25+
- Docker и Docker Compose
- Node.js 20+

## Архитектура

Проект состоит из двух частей:

- **frontend** — клиентская часть (Telegram Web App, Vue 3)
- **microservices** — два независимых бэкенд-сервиса:
  - `user_service` — управление пользователями и подписками
  - `wishlist_service` — управление списками желаний

Каждый сервис имеет собственную базу данных PostgreSQL.

Взаимодействие между сервисами — **gRPC**.
Связь frontend с backend — **REST API** через **gRPC-Gateway**.

## Стек

| Слой | Технологии |
|---|---|
| Frontend | Vue 3, Vite |
| Backend | Go, gRPC, gRPC-Gateway |
| База данных | PostgreSQL |
| Инфраструктура | Docker |

## Возможности

- Регистрация через Telegram
- Подписка на других пользователей
- Создание и просмотр вишлиста
- Добавление желаний с фото, ценой и ссылкой

> В разработке: удаление желаний, статусы желаний, уведомления

## Запуск

### 1. Настрой переменные окружения

```bash
cp microservices/user_service/.env.example microservices/user_service/.env
cp microservices/wishlist_service/.env.example microservices/wishlist_service/.env
```

Открой каждый `.env` и заполни значения по образцу из `.env.example`.

### 2. Подними базы данных

```bash
cd microservices/user_service
docker compose up -d

cd ../wishlist_service
docker compose up -d
```

### 3. Запусти сервисы

```bash
# user_service
cd microservices/user_service
go run cmd/main/main.go

# wishlist_service (в отдельном терминале)
cd microservices/wishlist_service
go run cmd/main/main.go
```

### 4. Запусти frontend

```bash
cd frontend
npm install
npm run dev
```

## Порты

| Сервис | gRPC | HTTP |
|---|---|---|
| user_service | :3030 | :5555 |
| wishlist_service | :8080 | :8081 |
| frontend | — | :5173 |
