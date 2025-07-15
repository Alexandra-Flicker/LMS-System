# 1. Базовый образ
FROM golang:1.24-alpine AS builder

# 2. Установка зависимостей
RUN apk add --no-cache git

# 3. Рабочая директория внутри контейнера
WORKDIR /app

# 4. Копируем go.mod и go.sum
COPY go.mod go.sum ./
RUN go mod download

# 5. Копируем остальной исходный код
COPY . .

# 6. Установка goose CLI (если надо использовать в миграциях)
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

# 7. Сборка приложения
RUN go build -o main ./cmd/lms

# 8. Финальный контейнер
FROM alpine:latest

# 9. Установка зависимостей (только для runtime)
RUN apk --no-cache add ca-certificates

# 10. Рабочая директория
WORKDIR /root/

# 11. Копируем собранное приложение
COPY --from=builder /app/main .

# 12. Копируем миграции, если они нужны
COPY --from=builder /app/migration ./migration

# 13. Порт, который слушает приложение
EXPOSE 8080

# 14. Команда запуска
CMD ["./main"]
