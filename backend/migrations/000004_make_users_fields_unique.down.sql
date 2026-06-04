ALTER TABLE IF EXISTS affiliate_system.users
    DROP CONSTRAINT IF EXISTS users_email_key;

ALTER TABLE IF EXISTS affiliate_system.users
    DROP CONSTRAINT IF EXISTS users_password_hash_key;