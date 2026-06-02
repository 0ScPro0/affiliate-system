CREATE TABLE IF NOT EXISTS affiliate_system.city (
    id          SERIAL        PRIMARY KEY,
    name        VARCHAR(50)  NOT NULL,
    CHECK (char_length(name) BETWEEN 1 AND 50)
);

ALTER TABLE affiliate_system.offer
ADD COLUMN city_id INTEGER NOT NULL;

ALTER TABLE affiliate_system.offer
ADD CONSTRAINT fk_offer_city 
    FOREIGN KEY (city_id) REFERENCES affiliate_system.city(id) ON DELETE RESTRICT;