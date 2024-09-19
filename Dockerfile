# Стадия 1: сборка
FROM golang:1.20-alpine AS builder

# Устанавливаем зависимости для сборки
RUN apk add --no-cache git

# Создаем рабочую директорию
WORKDIR /app

# Копируем файлы go.mod и go.sum для установки зависимостей
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем остальные файлы проекта
COPY . .

# Собираем бинарный файл с флагами оптимизации
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /todo-app ./cmd/main.go

# Стадия 2: финальный контейнер
FROM alpine:latest

# Создаем пользователя для безопасного запуска
RUN adduser -D -u 1000 todoapp

# Устанавливаем директорию для приложения
WORKDIR /home/todoapp

# Копируем бинарник из стадии сборки
COPY --from=builder /todo-app .

# Меняем права на исполняемый файл
RUN chown -R todoapp:todoapp /home/todoapp

# Переходим в пользователя с ограниченными правами
USER todoapp

# Экспортируем порт для доступа к приложению
EXPOSE 8000

# Запуск приложения
CMD ["./todo-app"]
