//запуск микросервиса
CONFIG_PATH=./config/local.yaml go run ./cmd/chiapp/main.go

go run ./cmd/ginapp/main.go

go run ./cmd/docsapp

структура проекта:
стандартная,
папка pkg опциональная

## lib

### библиотеки для файлов конфигураций

- cleanenv

### библиотеки для логирования

- slog

### db

- sqLite
- postgresql

### router

- chi, "chi render"

<br>

## ? вопросики на знание го:

что такое "struct tag"
constraint

## соглашения:

1. называть функции yначиная с Must... если функция будет паниковать а не возвращать ошибку (MustLoad) (использовать в крайних случаях)
