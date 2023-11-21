# Используем образ Golang
FROM golang:1.17 AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы проекта в рабочую директорию
COPY . .

# Собираем приложение
RUN go build -o main .

# Финальный образ
FROM alpine:latest

# Устанавливаем зависимости
RUN apk --no-cache add ca-certificates

# Копируем бинарник из предыдущего образа
COPY --from=builder /app/main /usr/local/bin/main

# Устанавливаем рабочую директорию
WORKDIR /usr/local/bin

# Команда для запуска приложения
CMD ["main"]
