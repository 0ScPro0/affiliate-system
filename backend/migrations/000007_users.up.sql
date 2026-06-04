CREATE TABLE IF NOT EXISTS affiliate_system.users (
    id        SERIAL        PRIMARY KEY,
    username  VARCHAR(50)   NOT NULL,
    email     VARCHAR(100)  NOT NULL,
    CHECK (char_length(username) BETWEEN 1 AND 50),
    CHECK (char_length(email) BETWEEN 1 AND 100)
)

