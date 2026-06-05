ALTER TABLE IF EXISTS affiliate_system.users
    DROP COLUMN IF EXISTS refresh_token,
    DROP COLUMN IF EXISTS refresh_token_expires_at;