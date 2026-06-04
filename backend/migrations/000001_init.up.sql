CREATE SCHEMA IF NOT EXISTS affiliate_system;

CREATE TABLE IF NOT EXISTS affiliate_system.users (
    id             SERIAL        PRIMARY KEY,
    username       VARCHAR(50)   NOT NULL,
    email          VARCHAR(100)  NOT NULL,
    password_hash  VARCHAR(255)  NOT NULL,
    is_admin       BOOLEAN       DEFAULT FALSE,
    created_at     TIMESTAMPTZ   NOT NULL  DEFAULT CURRENT_TIMESTAMP,
    CHECK (char_length(username) BETWEEN 1 AND 50),
    CHECK (char_length(email) BETWEEN 1 AND 100)
);

CREATE TABLE IF NOT EXISTS affiliate_system.partners (
    id          SERIAL        PRIMARY KEY,
    name        VARCHAR(100)  NOT NULL,
    description VARCHAR(1000),
    created_at  TIMESTAMPTZ   NOT NULL  DEFAULT CURRENT_TIMESTAMP,
    CHECK (char_length(name) BETWEEN 1 AND 100),
    CHECK (description IS NULL OR char_length(description) BETWEEN 1 AND 1000)
);

CREATE TABLE IF NOT EXISTS affiliate_system.categories (
    id          SERIAL        PRIMARY KEY,
    name        VARCHAR(100)  NOT NULL,
    description VARCHAR(1000),
    created_at  TIMESTAMPTZ   NOT NULL  DEFAULT CURRENT_TIMESTAMP,
    CHECK (char_length(name) BETWEEN 1 AND 100),
    CHECK (description IS NULL OR char_length(description) BETWEEN 1 AND 1000)
);

CREATE TABLE IF NOT EXISTS affiliate_system.cities (
    id          SERIAL        PRIMARY KEY,
    name        VARCHAR(50)  NOT NULL,
    created_at  TIMESTAMPTZ   NOT NULL  DEFAULT CURRENT_TIMESTAMP,
    CHECK (char_length(name) BETWEEN 1 AND 50)
);

CREATE TABLE IF NOT EXISTS affiliate_system.offers (
    id          SERIAL        PRIMARY KEY,
    partner_id  INTEGER       NOT NULL,
    category_id INTEGER       NOT NULL,
    city_id     INTEGER       NOT NULL,
    name        VARCHAR(100)  NOT NULL,
    description VARCHAR(1000),
    created_at  TIMESTAMPTZ   NOT NULL  DEFAULT CURRENT_TIMESTAMP,
    expire_at   TIMESTAMPTZ   NOT NULL,
    CHECK (char_length(name) BETWEEN 1 AND 100),
    CHECK (description IS NULL OR char_length(description) BETWEEN 1 AND 1000),
    FOREIGN KEY (partner_id) REFERENCES affiliate_system.partners(id) ON DELETE RESTRICT,
    FOREIGN KEY (category_id) REFERENCES affiliate_system.categories(id) ON DELETE RESTRICT,
    FOREIGN KEY (city_id) REFERENCES affiliate_system.cities(id) ON DELETE RESTRICT
);

CREATE INDEX IF NOT EXISTS idx_offer_partner_id ON affiliate_system.offers(partner_id);
CREATE INDEX IF NOT EXISTS idx_offer_category_id ON affiliate_system.offers(category_id);
CREATE INDEX IF NOT EXISTS idx_offer_city_id ON affiliate_system.offers(city_id);
CREATE INDEX IF NOT EXISTS idx_offer_expire_at ON affiliate_system.offers(expire_at);