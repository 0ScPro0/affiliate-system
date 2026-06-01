CREATE SCHEMA IF NOT EXISTS affiliate_system;

CREATE TABLE IF NOT EXISTS affiliate_system.partner (
    id          SERIAL        PRIMARY KEY,
    name        VARCHAR(100)  NOT NULL,
    description VARCHAR(1000),
    CHECK (char_length(name) BETWEEN 1 AND 100),
    CHECK (description IS NULL OR char_length(description) BETWEEN 1 AND 1000)
);

CREATE TABLE IF NOT EXISTS affiliate_system.category (
    id          SERIAL        PRIMARY KEY,
    name        VARCHAR(100)  NOT NULL,
    description VARCHAR(1000),
    CHECK (char_length(name) BETWEEN 1 AND 100),
    CHECK (description IS NULL OR char_length(description) BETWEEN 1 AND 1000)
);

CREATE TABLE IF NOT EXISTS affiliate_system.offer (
    id          SERIAL        PRIMARY KEY,
    partner_id  INTEGER       NOT NULL,
    category_id INTEGER       NOT NULL,
    name        VARCHAR(100)  NOT NULL,
    description VARCHAR(1000),
    created_at  TIMESTAMPTZ   NOT NULL  DEFAULT CURRENT_TIMESTAMP,
    expire_at   TIMESTAMPTZ   NOT NULL,
    CHECK (char_length(name) BETWEEN 1 AND 100),
    CHECK (description IS NULL OR char_length(description) BETWEEN 1 AND 1000),
    FOREIGN KEY (partner_id) REFERENCES affiliate_system.partner(id) ON DELETE RESTRICT,
    FOREIGN KEY (category_id) REFERENCES affiliate_system.category(id) ON DELETE RESTRICT
);

CREATE INDEX IF NOT EXISTS idx_offer_partner_id ON affiliate_system.offer(partner_id);
CREATE INDEX IF NOT EXISTS idx_offer_category_id ON affiliate_system.offer(category_id);
CREATE INDEX IF NOT EXISTS idx_offer_expire_at ON affiliate_system.offer(expire_at);