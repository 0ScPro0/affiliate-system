ALTER TABLE IF EXISTS affiliate_system.users
    ADD CONSTRAINT users_email_key UNIQUE (email);

ALTER TABLE IF EXISTS affiliate_system.users
    ADD CONSTRAINT users_password_hash_key UNIQUE (password_hash);