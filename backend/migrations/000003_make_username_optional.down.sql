ALTER TABLE IF EXISTS affiliate_system.users
    ALTER COLUMN username SET NOT NULL;

ALTER TABLE IF EXISTS affiliate_system.users
    ADD CONSTRAINT users_username_check CHECK (char_length(username) BETWEEN 1 AND 50);