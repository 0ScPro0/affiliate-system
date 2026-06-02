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
cmd/
├── server/
│   └── main.go                 # HTTP сервер (CRUD + API для событий)
└── consumer/
    └── main.go                 # Kafka consumer (отдельный процесс)

internal/
├── core/                       # Общая инфраструктура
│   ├── config/                 # Конфигурация (env + yaml)
│   ├── database/               # Подключение к Postgres (pgx/sqlx)
│   ├── domain/                 # Общие доменные типы (если есть)
│   ├── errors/                 # Кастомные ошибки
│   ├── logger/                 # Логирование (slog/logrus)
│   └── transport/
│       ├── http/               # HTTP клиент/сервер (если нужен)
│       │   ├── middleware/     # CORS, logging, recovery
│       │   └── response/       # Стандартные ответы API
│       └── kafka/              # Общие Kafka компоненты
│           ├── consumer.go     # Базовый consumer (обёртка)
│           ├── producer.go     # Базовый producer (на всякий случай)
│           └── types.go        # Общие типы для Kafka сообщений

├── features/                   # Бизнес-фичи
│   ├── affiliate/              # Основная фича (партнёры, категории, офферы)
│   │   ├── service/            # Бизнес-логика
│   │   │   ├── partner.go
│   │   │   ├── category.go
│   │   │   └── offer.go
│   │   ├── repository/         # Работа с БД (CRUD)
│   │   │   ├── partner.go
│   │   │   ├── category.go
│   │   │   └── offer.go
│   │   ├── handler/            # HTTP handlers
│   │   │   ├── partner.go
│   │   │   ├── category.go
│   │   │   └── offer.go
│   │   ├── models/             # DTO/Entity (можно вынести в domain)
│   │   │   ├── partner.go
│   │   │   ├── category.go
│   │   │   └── offer.go
│   │   └── routes.go           # Регистрация роутов для этой фичи
│   │
│   └── events/                 # Для работы с событиями из Kafka
│       ├── service/            # Бизнес-логика обработки событий
│       │   └── event_processor.go
│       ├── repository/         # Работа с event_log таблицей
│       │   └── event.go
│       ├── handler/            # HTTP handlers для отдачи событий UI
│       │   └── event.go
│       ├── consumer/           # Специфичный consumer для этой фичи
│       │   └── wal_handler.go  # Обработка сообщений от wal-listener
│       ├── models/             # Модели событий
│       │   └── event.go
│       └── routes.go           # Регистрация /api/events

pkg/                            # Переиспользуемые пакеты (можно вынести в отдельные репозитории)
├── kafka/                      # Если нужно вынести общую логику Kafka
└── postgres/                   # Хелперы для Postgres

migrations/                     # SQL миграции
├── 001_create_schemas.up.sql
├── 001_create_schemas.down.sql
├── 002_create_event_log.up.sql
└── 002_create_event_log.down.sql

configs/
├── config.yaml                 # Основной конфиг
├── wal-listener.yaml           # Конфиг для wal-listener
└── .env.example                # Пример env переменных

docker-compose.yml
Makefile                        # Упрощение команд (make up, make migrate, etc.)
go.mod
go.sum
README.md
```



### 2. Схема БД:

```sql
-- Предложение (offer):
TABLE offer(
    id           SERIAL                   PRIMARY KEY,
    partner_id                            FOREIGN KEY,
    category_id                           FOREIGN KEY,
    name         VARCHAR(100)   NOT NULL  CHECK (char_length(name) BETWEEN 1 AND 100),
    description  VARCHAR(1000)            CHECK (char_length(description) BETWEEN 1 AND 1000), 
    created_at   TIMESTAMPTZ    NOT NULL  DEFAULT CURRENT_TIMESTAMP,         
    expire_at    TIMESTAMPTZ    NOT NULL            
)

-- Партнер (partner):
TABLE partner(
    id           SERIAL                   PRIMARY KEY,
    name         VARCHAR(100)   NOT NULL  CHECK (char_length(name) BETWEEN 1 AND 100),
    description  VARCHAR(1000)            CHECK (char_length(description) BETWEEN 1 AND 1000)    
)

-- Категория (category):
TABLE category(
    id           SERIAL                   PRIMARY KEY,
    name         VARCHAR(100)   NOT NULL  CHECK (char_length(name) BETWEEN 1 AND 100),
    description  VARCHAR(1000)            CHECK (char_length(description) BETWEEN 1 AND 1000) 
)
```