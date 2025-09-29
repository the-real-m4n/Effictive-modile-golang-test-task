# Сервис для Управления Подписками

Этот проект представляет собой веб-приложение для управления подписками, разработанное с использованием Go на бэкенде и React на фронтенде.

## Возможности

- **CRUD операции:** Создание, чтение, обновление и удаление подписок.
- **Подсчет стоимости:** Автоматический расчет суммарной стоимости активных подписок за выбранный или текущий месяц.
- **Единый сервис:** Фронтенд и бэкенд работают в одном Docker-контейнере для простоты запуска и развертывания.

## Технологии

- **Бэкенд:** Go, Gin, PostgreSQL
- **Фронтенд:** React, React Bootstrap, React Datepicker
- **Окружение:** Docker, Docker Compose

## Предварительные требования

Перед запуском убедитесь, что у вас установлены:

- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)
- `make` (обычно предустановлен в Linux/macOS, для Windows можно установить через [Chocolatey](https://chocolatey.org/install) или [WSL](https://docs.microsoft.com/en-us/windows/wsl/install))
- `golang-migrate` ([инструкция по установке](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate#installation))

## Запуск проекта

1. **Запустите Docker-контейнеры:**
   Эта команда соберет образ и запустит бэкенд вместе с базой данных в фоновом режиме.
   ```bash
   docker-compose up -d --build
   ```

2. **Примените миграции базы данных:**
   Эта команда создаст необходимую таблицу `subscriptions` в базе данных.
   ```bash
   make migrate-up
   ```

3. **Откройте приложение:**
   Теперь приложение доступно в вашем браузере по адресу:
   [http://localhost:8080](http://localhost:8080)

## API Эндпоинты

- `GET /subscriptions`: Получить список всех подписок.
- `POST /subscriptions`: Создать новую подписку.
- `PUT /subscriptions/{id}`: Обновить подписку по ID.
- `DELETE /subscriptions/{id}`: Удалить подписку по ID.
- `GET /subscriptions/monthly-cost`: Получить суммарную стоимость за месяц (можно передать query-параметр `month=YYYY-MM`).
- `GET /swagger/*any`: Документация Swagger UI.
