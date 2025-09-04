

# 📌 Subscriptions Service

![Go](https://img.shields.io/badge/Go-1.24-blue)
![Postgres](https://img.shields.io/badge/Postgres-15-blueviolet)
![License](https://img.shields.io/badge/license-MIT-green)

## 📖 Описание

`Subscriptions Service` — это **REST API на Go** для управления пользовательскими подписками.

Функционал:

* 📦 CRUD-операции с подписками
* 💰 Подсчёт суммарной стоимости подписок за период
* 📑 Автоматическая Swagger-документация
* 🐳 Упаковка в Docker (API + БД)

---

## ⚙️ Стек

* **Go 1.24** + [Gin](https://github.com/gin-gonic/gin)
* **PostgreSQL 15** + [pgx](https://github.com/jackc/pgx)
* **Swagger (swaggo)** для автогенерации документации
* **Docker & docker-compose**

---

## 📂 Структура проекта

```
.
├── cmd/              # main.go (точка входа)
├── internal/
│   ├── db/               # подключение к БД
│   ├── handler/          # HTTP-хендлеры
│   ├── models/           # модели
│   ├── repository/       # SQL-запросы
│   └── service/          # бизнес-логика (опц.)
├── migrations/           # SQL-миграции
├── docs/                 # Swagger-документация
├── Dockerfile
├── docker-compose.yml
└── README.md
```

---

## 🚀 Запуск проекта

### 1. Клонирование

```bash
git clone https://github.com/the-real-m4n/Effictive-modile-golang-test-task.git
cd subscriptions-service
```

### 2. Запуск в Docker

```bash
docker compose up --build
```

### 3. Доступы

* **API**: [http://localhost:8080](http://localhost:8080)
* **Swagger UI**: [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)
* **PostgreSQL**: `localhost:5433`

  * user: `postgres`
  * password: `postgres`
  * db: `subscriptions`

---

## 📖 API (Примеры)

### ➕ Создать подписку

```http
POST /subscriptions
Content-Type: application/json

{
  "service_name": "Yandex Plus",
  "price": 400,
  "user_id": "60601fee-2bf1-4721-ae6f-7636e79a0cba",
  "start_date": "2025-07",
  "end_date": "2025-12"
}
```

### 📜 Получить все подписки

```http
GET /subscriptions
```

### 🔎 Получить по ID

```http
GET /subscriptions/1
```

### ✏️ Обновить

```http
PUT /subscriptions/1
Content-Type: application/json

{
  "service_name": "Spotify",
  "price": 600,
  "start_date": "2025-08",
  "end_date": "2025-12"
}
```

### ❌ Удалить

```http
DELETE /subscriptions/1
```

### 💰 Посчитать сумму

```http
GET /subscriptions/total?user_id=60601fee-2bf1-4721-ae6f-7636e79a0cba&service_name=Yandex Plus&from=2025-07&to=2025-12
```

**Ответ**

```json
{
  "total": 2400
}
```

---

## 🗄️ Миграции

```bash
migrate -path ./migrations -database "postgres://postgres:postgres@localhost:5433/subscriptions?sslmode=disable" up
```
