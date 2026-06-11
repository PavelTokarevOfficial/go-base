# Go Base

Учебный проект на Go с PostgreSQL в Docker.

## Запуск

```bash
docker compose up -d
```

Проверить подключение Go-приложения:

```bash
go run main.go
```

## Админка БД

Pgweb доступен в браузере:

http://localhost:8080

# Миграции
## Сделать миграции
```bash
migrate \
  -path migrations \
  -database "postgres://go_user:go_pass@127.0.0.1:5433/go_db?sslmode=disable" \
  up
```

## Откатить миграции
```bash
migrate \
  -path migrations \
  -database "postgres://go_user:go_pass@127.0.0.1:5433/go_db?sslmode=disable" \
  down
```