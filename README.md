# Metric collection service 

Сервис предназначен для сбора телеметрии с мобильных приложений с целью последующей аналитики. 
Представляет собой HTTP-сервер с одним обработчиком POST-запросов.
## Миграции

### Запуск миграции

```make local-migration-up```

Запускает миграцию базы данных вверх с использованием goose

### Откат локальной миграции

```make local-migration-down```

Откатывает локальную миграцию базы данных с использованием goose

### Статус локальной миграции

```make local-migration-status```

Отображает статус локальной миграции базы данных с использованием goose

## Docker Compose

### Запуск Docker Compose

```make docker-compose-up ```

Запускает Docker Compose в фоновом режиме. Создает database postgres

## Запуск приложения

### Запуск приложения

```make run```

Запускает приложение с указанием пути к конфигурационному файлу и параметрами подключения к базе данных

## Тестирование

### Запуск тестов

```make test```

Запускает все тесты

### Запуск тестов с покрытием

```make test-coverage```

Запускает тесты с генерацией отчета о покрытии кода

### Запуск тестов с подробным выводом

```make test-v-cover```

Запускает тесты с подробным выводом

## Линтер

### Запуск линтера

```make lint```

Запускает golangci-lint
