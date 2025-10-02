# 1. Билд-стейдж только для установки зависимостей и тулов
FROM golang:1.25 AS base

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

# Установим air (live reload) и goose (миграции)
RUN go install github.com/air-verse/air@latest \
 && go install github.com/pressly/goose/v3/cmd/goose@latest

# 2. Runtime-стейдж (тоже golang, чтобы был компилятор под air)
FROM golang:1.25

WORKDIR /app
ENV PATH="/go/bin:${PATH}"

# Перенесём скачанные модули, чтобы ускорить повторные билды
COPY --from=base /go/pkg /go/pkg
COPY --from=base /go/bin /go/bin

# Копируем всё приложение
COPY . .

# Добавим entrypoint
COPY docker/dev/app/entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

# Откроем порт приложения (если у тебя :8080)
EXPOSE 8080

CMD ["/entrypoint.sh"]