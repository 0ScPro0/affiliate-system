ALTER TABLE IF EXISTS affiliate_system.users
    ADD COLUMN IF NOT EXISTS refresh_token VARCHAR(512) NULL,
    ADD COLUMN IF NOT EXISTS refresh_token_expires_at TIMESTAMPTZ NULL;