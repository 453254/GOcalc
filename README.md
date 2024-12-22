# Веб-сервис Калькулятор

Простой веб-сервис для вычисления арифметических выражений. Сервис принимает строковые математические выражения и возвращает результат их вычисления. Поддерживаются основные арифметические операции: сложение, вычитание, умножение и деление.

## Установка

1. Клонируйте репозиторий:
```bash
git clone https://github.com/your-username/calculator-service.git
cd calculator-service
```

2. Установите зависимости:
```bash
go mod download
```

## Запуск сервиса

Чтобы запустить сервис, выполните:

```bash
go run ./cmd/calc_service/...
```

Сервис запустится на порту 8080.

## Структура проекта

```
.
├── cmd/
│   └── calc_service/        # Точка входа в приложение
├── internal/
│   ├── api/                # Обработчики API
│   └── calculator/         # Логика вычислений
├── go.mod
├── go.sum
└── README.md
```

## Использование API

### Вычисление выражения

Отправьте POST запрос на `/api/v1/calculate` с JSON-телом, содержащим выражение для вычисления.

**Формат запроса:**
```bash
curl --location 'localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
    "expression": "2+2*2"
}'
```

**Успешный ответ (200):**
```json
{
    "result": "6"
}
```

**Неверное выражение (422):**
```json
{
    "error": "Expression is not valid"
}
```

**Ошибка сервера (500):**
```json
{
    "error": "Internal server error"
}
```

### Примеры использования

**1. Успешное вычисление:**
```bash
curl --location 'localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
    "expression": "2+2*2"
}'

# Ответ:
# {
#     "result": "6"
# }
```

**2. Неверное выражение:**
```bash
curl --location 'localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
    "expression": "2++2"
}'

# Ответ:
# {
#     "error": "Expression is not valid"
# }
```

**3. Деление на ноль (ошибка 500):**
```bash
curl --location 'localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
    "expression": "1/0"
}'

# Ответ:
# {
#     "error": "Internal server error"
# }
```

## Требования

- Go 1.16 или выше
- пакет gorilla/mux (`go get github.com/gorilla/mux`)
