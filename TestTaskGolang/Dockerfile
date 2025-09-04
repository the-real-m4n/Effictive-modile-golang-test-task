# Этап сборки
FROM golang:1.24 AS builder

WORKDIR /TestTaskGolang

# копируем go.mod и go.sum для кеша зависимостей
COPY go.mod go.sum ./
RUN go mod download

# копируем весь проект
COPY . .

# собираем бинарник
RUN go build -o subscriptions-service ./cmd

# Этап запуска
FROM debian:12-slim

WORKDIR /TestTaskGolang

# скопируем бинарник
COPY --from=builder /TestTaskGolang/subscriptions-service .

# swagger файлы (если нужны в контейнере)
COPY --from=builder /TestTaskGolang/docs ./docs

# порт, на котором работает gin
EXPOSE 8080

CMD ["./subscriptions-service"]
