ALTER TABLE IF EXISTS affiliate_system.users
    ALTER COLUMN username DROP NOT NULL,
    DROP CONSTRAINT IF EXISTS users_username_check;