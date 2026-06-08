[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/ihippik/wal-listener)
[![Vue.js](https://img.shields.io/badge/Vue.js-4FC08D?logo=vue.js&logoColor=white)](https://vuejs.org/)
[![Kafka](https://img.shields.io/badge/Kafka-231F20?logo=apachekafka&logoColor=white)](https://kafka.apache.org/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-4169E1?logo=postgresql&logoColor=white)](https://www.postgresql.org/)
[![WAL-Listener](https://img.shields.io/badge/WAL--Listener-181717?logo=github&logoColor=white)](https://github.com/ihippik/wal-listener)

**Реализация партнерской программы по ТЗ**

## ТЗ партнерской программы:

1. Реализовать партнерскую систему. Пример (https://my.ufanet.ru/partners/sales). Нужно
сделать разделы, карточки. Требуется сделать UI для отображения результата. Проект
должен иметь БД в которую будет складывать информацию об карточках и разделах.
В качестве БД используют Postgres.
2. Нужно подключится к БД с помощью https://github.com/ihippik/wal-listener и считывать
события из БД и класть их в топик Kafka.
3. Создать слушатель Kafka и сделать или телеграмм бот или UI на котором выводятся
события об изменениях в БД.
В итоге должна получится цепочка. Пользователь через админку партнеров сохраняет
в БД-> Перехватчик слушает события с БД и кладет их в Kafka-> Слушатель кафки
выводит событие в интерфейс.
4. Обложить все тестами. Coverage 80% минимум.

## Реализация

### 1. Архитектура бекэнда:

```bash
backend/
├── cmd/
│   ├── server/
│   │   └── main.go                 # HTTP сервер (точка входа)
│   └── consumer/
│       └── main.go                 # Kafka consumer (отдельный процесс, заглушка)
│
├── internal/
│   ├── core/                       # Общая инфраструктура
│   │   ├── config/
│   │   │   ├── config.go           # Конфигурация (env variables)
│   │   │   └── wal-config.yml      # Конфигурация для WAL-listener
│   │   │
│   │   ├── database/
│   │   │   ├── models/             # Модели БД (CategoryModel, CityModel, OfferModel, PartnerModel, UserModel)
│   │   │   └── postgres/pool/
│   │   │       └── pool.go         # Пул соединений к PostgreSQL (pgx)
│   │   │
│   │   ├── domain/
│   │   │   ├── entity/             # Доменные сущности (Category, City, Offer, Partner, User)
│   │   │   └── event/              # Доменные события (DomainEvent interface)
│   │   │       ├── events.go       # Базовый интерфейс и структура событий
│   │   │       ├── category/       # События категорий (created, updated, deleted)
│   │   │       ├── city/           # События городов (created, updated, deleted)
│   │   │       ├── offer/          # События предложений (created, updated, deleted)
│   │   │       ├── partner/        # События партнеров (created, updated, deleted)
│   │   │       └── user/           # События пользователей (created, updated, deleted)
│   │   │
│   │   ├── errors/
│   │   │   └── common.go           # Кастомные ошибки (ErrNotFound, ErrInvalidArgument, ErrConflict, ErrUnauthorized, ErrForbidden)
│   │   │
│   │   ├── logger/
│   │   │   └── logger.go           # Логирование (zap, console + file)
│   │   │
│   │   ├── security/
│   │   │   ├── jwt.go              # JWT токены (access + refresh)
│   │   │   └── password.go         # Хеширование паролей (bcrypt)
│   │   │
│   │   ├── transport/
│   │   │   ├── dto/                # DTO для запросов/ответов (auth, category, city, offer, partner, user)
│   │   │   └── http/
│   │   │       ├── middleware/
│   │   │       │   ├── middleware.go   # Middleware chain
│   │   │       │   ├── auth.go         # JWT аутентификация
│   │   │       │   └── common.go       # CORS, RequestID, Logger, Panic, Trace
│   │   │       ├── request/
│   │   │       │   └── decode.go       # Декодирование и валидация JSON запросов
│   │   │       ├── response/
│   │   │       │   ├── handler.go      # Обработчик HTTP ответов
│   │   │       │   ├── writer.go       # Обертка ResponseWriter
│   │   │       │   └── errors.go       # Структура ErrorResponse
│   │   │       └── server/
│   │   │           ├── server.go       # HTTP сервер (запуск, graceful shutdown)
│   │   │           ├── router.go       # APIVersionRouter (v1, v2, v3)
│   │   │           └── route.go        # Структура Route
│   │   │
│   │   └── utils/
│   │       ├── string.go           # Утилиты для строк
│   │       └── validate.go         # Валидация длины строк
│   │
│   └── features/                   # Бизнес-фичи
│       └── affiliate/
│           ├── auth/               # Аутентификация (register, login, logout, refresh_token)
│           │   ├── service/
│           │   └── transport/http/
│           │
│           ├── category/           # CRUD категорий
│           │   ├── repository/postgres/
│           │   ├── service/
│           │   └── transport/http/
│           │
│           ├── city/               # CRUD городов
│           │   ├── repository/postgres/
│           │   ├── service/
│           │   └── transport/http/
│           │
│           ├── offer/              # CRUD предложений (репозиторий + сервис есть, HTTP транспорт отсутствует)
│           │   ├── repository/postgres/
│           │   ├── service/
│           │   └── transport/      # пусто
│           │
│           ├── partner/            # CRUD партнеров
│           │   ├── repository/postgres/
│           │   ├── service/
│           │   └── transport/http/
│           │
│           └── user/               # CRUD пользователей
│               ├── repository/postgres/
│               ├── service/
│               └── transport/http/
│
├── docs/                           # Swagger документация (автогенерация)
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
│
├── migrations/                     # SQL миграции
│   ├── 000001_init.up/down.sql         # Создание схемы и таблиц (users, partners, categories, cities, offers)
│   ├── 000002_db_event.up/down.sql     # Таблица event_log для отслеживания изменений
│   ├── 000003_make_username_optional.up/down.sql  # username стал опциональным
│   ├── 000004_make_users_fields_unique.up/down.sql # UNIQUE constraints на email и password_hash
│   └── 000005_refresh_tokens.up/down.sql           # Добавление refresh_token полей
│
├── go.mod
├── go.sum
│
└── out/                            # Выходные директории (создаются во время работы)
    ├── logs/                       # Лог-файлы
    ├── pgdata/                     # Данные PostgreSQL (Docker volume)
    └── kafka_data/                 # Данные Kafka (Docker volume)
```

### 2. Схема БД:

```sql
-- Схема: affiliate_system

-- Пользователь (user):
TABLE users(
    id                    SERIAL        PRIMARY KEY,
    username              VARCHAR(50),                          -- опционально (был NOT NULL, изменено в миграции 000003)
    email                 VARCHAR(100)  NOT NULL  UNIQUE,       -- UNIQUE добавлено в миграции 000004
    password_hash         VARCHAR(255)  NOT NULL  UNIQUE,       -- UNIQUE добавлено в миграции 000004
    is_admin              BOOLEAN       DEFAULT FALSE,
    refresh_token         VARCHAR(512),                         -- добавлено в миграции 000005
    refresh_token_expires_at TIMESTAMPTZ,                       -- добавлено в миграции 000005
    created_at            TIMESTAMPTZ   NOT NULL  DEFAULT CURRENT_TIMESTAMP,
    CHECK (char_length(email) BETWEEN 1 AND 100)
)

-- Партнер (partner):
TABLE partners(
    id          SERIAL        PRIMARY KEY,
    name        VARCHAR(100)  NOT NULL,
    description VARCHAR(1000),
    created_at  TIMESTAMPTZ   NOT NULL  DEFAULT CURRENT_TIMESTAMP,
    CHECK (char_length(name) BETWEEN 1 AND 100),
    CHECK (description IS NULL OR char_length(description) BETWEEN 1 AND 1000)
)

-- Категория (category):
TABLE categories(
    id          SERIAL        PRIMARY KEY,
    name        VARCHAR(100)  NOT NULL,
    description VARCHAR(1000),
    created_at  TIMESTAMPTZ   NOT NULL  DEFAULT CURRENT_TIMESTAMP,
    CHECK (char_length(name) BETWEEN 1 AND 100),
    CHECK (description IS NULL OR char_length(description) BETWEEN 1 AND 1000)
)

-- Город (city):
TABLE cities(
    id          SERIAL        PRIMARY KEY,
    name        VARCHAR(50)   NOT NULL,
    created_at  TIMESTAMPTZ   NOT NULL  DEFAULT CURRENT_TIMESTAMP,
    CHECK (char_length(name) BETWEEN 1 AND 50)
)

-- Предложение (offer):
TABLE offers(
    id          SERIAL        PRIMARY KEY,
    partner_id  INTEGER       NOT NULL  REFERENCES partners(id) ON DELETE RESTRICT,
    category_id INTEGER       NOT NULL  REFERENCES categories(id) ON DELETE RESTRICT,
    city_id     INTEGER       NOT NULL  REFERENCES cities(id) ON DELETE RESTRICT,
    name        VARCHAR(100)  NOT NULL,
    description VARCHAR(1000),
    created_at  TIMESTAMPTZ   NOT NULL  DEFAULT CURRENT_TIMESTAMP,
    expire_at   TIMESTAMPTZ   NOT NULL,
    CHECK (char_length(name) BETWEEN 1 AND 100),
    CHECK (description IS NULL OR char_length(description) BETWEEN 1 AND 1000)
)

-- Лог событий (event_log):
TABLE event_log(
    id          SERIAL        PRIMARY KEY,
    event_type  VARCHAR(50)   NOT NULL,
    entity_type VARCHAR(50)   NOT NULL,
    entity_id   INTEGER       NOT NULL,
    payload     JSONB         NOT NULL,
    created_at  TIMESTAMPTZ   NOT NULL  DEFAULT CURRENT_TIMESTAMP
)
```

### 3. API Endpoints:

Базовый путь: `/api/v1/`

| Метод      | Путь                        | Аутентификация | Описание                        |
|------------|-----------------------------|----------------|---------------------------------|
| **Auth**   |                             |                |                                 |
| `POST`     | `/register`                 | Нет            | Регистрация нового пользователя |
| `POST`     | `/login`                    | Нет            | Вход в систему                  |
| `POST`     | `/logout`                   | Да             | Выход из системы                |
| `POST`     | `/refresh_token`            | Нет            | Обновление access токена        |
| **Cities** |                             |                |                                 |
| `POST`     | `/cities`                   | Да             | Создание города                 |
| `GET`      | `/cities/{id}`              | Нет            | Получение города по ID          |
| `PATCH`    | `/cities/{id}`              | Да             | Обновление города               |
| `DELETE`   | `/cities/{id}`              | Да             | Удаление города                 |
| **Partners** |                           |                |                                 |
| `POST`     | `/partners`                 | Да             | Создание партнера               |
| `GET`      | `/partners/{id}`            | Нет            | Получение партнера по ID        |
| `PATCH`    | `/partners/{id}`            | Да             | Обновление партнера             |
| `DELETE`   | `/partners/{id}`            | Да             | Удаление партнера               |
| **Categories** |                        |                |                                 |
| `POST`     | `/categories`               | Да             | Создание категории              |
| `GET`      | `/categories/{id}`          | Нет            | Получение категории по ID       |
| `PATCH`    | `/categories/{id}`          | Да             | Обновление категории            |
| `DELETE`   | `/categories/{id}`          | Да             | Удаление категории              |
| **Users**  |                             |                |                                 |
| `POST`     | `/users`                    | Да             | Создание пользователя           |
| `GET`      | `/users/{id}`               | Да             | Получение пользователя по ID    |
| `PATCH`    | `/users/{id}`               | Да             | Обновление пользователя         |
| `DELETE`   | `/users/{id}`               | Да             | Удаление пользователя           |
| **Docs**   |                             |                |                                 |
| `GET`      | `/docs/*`                   | Нет            | Swagger UI документация         |

### 4. Технологический стек:

- **Язык:** Go 1.25
- **HTTP роутер:** Стандартный `net/http` + `http.ServeMux`
- **База данных:** PostgreSQL 18 с WAL (Logical Replication)
- **Брокер сообщений:** Apache Kafka 4.3
- **WAL-listener:** [ihippik/wal-listener](https://github.com/ihippik/wal-listener) v2.11.0
- **Миграции:** [golang-migrate/migrate](https://github.com/golang-migrate/migrate) v4.19.1
- **Логирование:** go.uber.org/zap
- **JWT:** golang-jwt/jwt/v5
- **Пароли:** bcrypt (golang.org/x/crypto)
- **Пул соединений:** jackc/pgx/v5 (pgxpool)
- **Валидация:** go-playground/validator/v10
- **Swagger:** swaggo/swag v1.16.6
- **Фронтенд:** Vue.js 2.x

### 5. Инфраструктура (Docker Compose):

Сервисы, запускаемые через `docker compose`:

| Сервис                          | Назначение                                    |
|---------------------------------|-----------------------------------------------|
| `affiliate-system-postgres`     | PostgreSQL с включенным WAL logical replication |
| `affiliate-system-kafka`        | Apache Kafka (KRaft mode, без Zookeeper)      |
| `affiliate-system-wal-listener` | Слушает WAL и публикует изменения в Kafka     |
| `affiliate-system-swagger`      | Генерация Swagger документации                |
| `affiliate-system-migrate`      | Выполнение SQL миграций БД                    |

### 6. Конфигурация (переменные окружения):

| Переменная                              | Описание                              |
|-----------------------------------------|---------------------------------------|
| `SERVER_HOST`                           | Хост HTTP сервера                     |
| `SERVER_PORT`                           | Порт HTTP сервера                     |
| `SERVER_SHUTDOWN_TIMEOUT`               | Таймаут graceful shutdown             |
| `DATABASE_POSTGRES_HOST`                | Хост PostgreSQL                       |
| `DATABASE_POSTGRES_PORT`                | Порт PostgreSQL                       |
| `DATABASE_POSTGRES_USER`                | Пользователь PostgreSQL               |
| `DATABASE_POSTGRES_PASSWORD`            | Пароль PostgreSQL                     |
| `DATABASE_POSTGRES_DB`                  | Название БД                           |
| `DATABASE_POSTGRES_URL`                 | Полный URL подключения к БД           |
| `DATABASE_POSTGRES_TIMEOUT`             | Таймаут операций с БД                 |
| `SECURITY_SECRET_KEY`                   | Секретный ключ для JWT                |
| `SECURITY_ACCESS_TOKEN_EXPIRE_MINUTES`  | Время жизни access токена (мин)       |
| `SECURITY_REFRESH_TOKEN_EXPIRE_DAYS`    | Время жизни refresh токена (дни)      |
| `LOGGER_LEVEL`                          | Уровень логирования                   |
| `LOGGER_FOLDER`                         | Папка для лог-файлов                  |
| `ENVIRONMENT_DEBUG`                     | Режим отладки                         |

### 7. Команды Makefile:

| Команда              | Описание                                          |
|----------------------|----------------------------------------------------|
| `make env-up`        | Запуск инфраструктуры (Postgres, Kafka, WAL-listener) |
| `make env-down`      | Остановка инфраструктуры                          |
| `make migrate-up`    | Применение миграций БД                            |
| `make migrate-down`  | Откат миграций БД                                 |
| `make migrate-status`| Статус миграций                                   |
| `make migrate-create seq=<name>` | Создание новой миграции                |
| `make run-server`    | Запуск HTTP сервера                               |
| `make run-consumer`  | Запуск Kafka consumer                             |
| `make swagger-gen`   | Генерация Swagger документации                    |

### 8. Архитектурные решения:

- **Чистая архитектура (Clean Architecture):** Код разделен на слои: транспорт (HTTP handlers) → сервис (бизнес-логика) → репозиторий (работа с БД).
- **Dependency Injection:** Зависимости (репозитории, сервисы) создаются в `main.go` и передаются через конструкторы.
- **Domain Events:** Каждая сущность генерирует события (created, updated, deleted), которые могут быть опубликованы в Kafka через WAL-listener.
- **JWT Authentication:** Двухтокенная аутентификация (access + refresh token) с middleware для защиты эндпоинтов.
- **Graceful Shutdown:** Сервер корректно обрабатывает сигналы SIGINT/SIGTERM.
- **Версионирование API:** Поддержка множества версий API через `APIVersionRouter` (v1, v2, v3).