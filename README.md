# TechSchool Project

Project with using Go, Kafka, PostgreSQL and Redis

## Requirements

- Go 1.24+
- PostgreSQL 15+
- Flyway CLI (опционально, если используете Docker)

## Установка Flyway

### macOS (с Homebrew)

```bash
brew install flyway
```

### Linux

```bash
# Скачайте Flyway с официального сайта
wget -qO- https://repo1.maven.org/maven2/org/flywaydb/flyway-commandline/9.22.3/flyway-commandline-9.22.3-linux-x64.tar.gz | tar xvz
sudo ln -s `pwd`/flyway-9.22.3/flyway /usr/local/bin/flyway
```

### Windows

Скачайте Flyway с [официального сайта](https://flywaydb.org/download) и добавьте в PATH.

## Настройка базы данных

1. Убедитесь, что PostgreSQL запущен
2. Создайте базу данных:

```bash
createdb techschool
```

## Конфигурация

Основные настройки Flyway находятся в файле `flyway.conf`:

- **URL базы данных**: `jdbc:postgresql://localhost:5432/techschool`
- **Пользователь**: `postgres`
- **Пароль**: `postgres`
- **Папка миграций**: `migrations/`
- **Префикс миграций**: `V`
- **Разделитель**: `__`
- **Суффикс**: `.sql`

## Использование

### Основные команды

```bash
# Применить все pending миграции
make migrate

# Показать статус миграций
make migrate-info

# Очистить базу данных (удалить все таблицы)
make migrate-clean

# Восстановить историю миграций
make migrate-repair

# Валидировать миграции
make migrate-validate

# Создать новую миграцию
make migrate-create
```

### Прямые команды Flyway

```bash
# Применить миграции
flyway -configFiles=flyway.conf migrate

# Показать информацию
flyway -configFiles=flyway.conf info

# Очистить базу
flyway -configFiles=flyway.conf clean

# Восстановить
flyway -configFiles=flyway.conf repair

# Валидировать
flyway -configFiles=flyway.conf validate
```

## Создание новых миграций

### Автоматическое создание

```bash
make migrate-create
# Введите описание миграции
```

### Ручное создание

Создайте файл в папке `migrations/` с именем в формате:

```
V{версия}__{описание}.sql
```

Пример: `V2__Add_user_table.sql`

## Структура миграций

```
migrations/
├── V1__Create_orders_tables.sql    # Создание таблиц заказов
└── V2__Add_user_table.sql          # Добавление таблицы пользователей
```

## Docker

### Запуск с Docker Compose

```bash
# Запустить все сервисы (PostgreSQL + Flyway + App)
docker-compose up -d

# Просмотр логов
docker-compose logs -f

# Остановить сервисы
docker-compose down
```

### Только база данных

```bash
# Запустить только PostgreSQL
docker-compose up postgres -d

# Применить миграции
make migrate
```

## Разработка

```bash
# Установить зависимости
make deps

# Запустить в режиме разработки
make dev

# Собрать приложение
make build

# Запустить приложение
make run

# Запустить тесты
make test
```

## Полезные команды

```bash
# Показать все доступные команды
make help

# Настройка базы данных
make db-setup

# Очистка артефактов сборки
make clean
```

## Troubleshooting

### Проблемы с подключением к базе данных

1. Убедитесь, что PostgreSQL запущен
2. Проверьте настройки в `flyway.conf`
3. Убедитесь, что база данных `techschool` существует

### Проблемы с миграциями

1. Проверьте синтаксис SQL в файлах миграций
2. Убедитесь, что имена файлов соответствуют конвенции Flyway
3. Используйте `make migrate-validate` для проверки

### Очистка и перезапуск

```bash
# Очистить базу данных
make migrate-clean

# Применить миграции заново
make migrate
```

## Структура проекта

```
.
├── migrations/           # SQL миграции
│   └── V1__Create_orders_tables.sql
├── models/              # Go модели
├── workerpool/          # Worker pool
├── flyway.conf          # Конфигурация Flyway
├── docker-compose.yml   # Docker конфигурация
├── Makefile            # Команды для разработки
├── go.mod              # Go модули
└── main.go             # Главный файл приложения
```
