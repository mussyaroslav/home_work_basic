# Этап 1: Сборка приложения
FROM golang:1.20 AS builder

WORKDIR /app

# Копируем все файлы в рабочую директорию контейнера
COPY . .

# Сборка Go-приложения
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# Этап 2: Финальный контейнер с приложением
FROM alpine:3.14

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем собранное приложение и конфигурацию
COPY --from=builder /app/main /app/main
COPY config.yml /app

# Даем права на исполнение
RUN chmod +x /app/main

# Открываем порт
EXPOSE 8080

# Запускаем приложение
CMD ["/app/main"]
