# Используем базовый образ с Go
FROM golang:1.22 AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum из корня проекта
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем всё содержимое проекта
COPY . .

# Переходим в директорию с main.go и собираем бинарный файл
WORKDIR ./cmd/
RUN go build -o /main

FROM ubuntu:22.04

WORKDIR /root/

# Копируем бинарный файл из этапа сборки
COPY --from=builder /main .
COPY ./web/ /root/web

# Копируем файл .env (если используется)
COPY .env .

# Запускаем приложение
CMD ["./main"]

