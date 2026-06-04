CREATE TABLE IF NOT EXISTS affiliate_system.event_log (
    id          SERIAL        PRIMARY KEY,
    event_type  VARCHAR(50)   NOT NULL,
    entity_type VARCHAR(50)   NOT NULL,
    entity_id   INTEGER       NOT NULL,
    payload     JSONB         NOT NULL,
    created_at  TIMESTAMPTZ   NOT NULL  DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_event_log_entity ON affiliate_system.event_log(entity_type, entity_id);
CREATE INDEX IF NOT EXISTS idx_event_log_created_at ON affiliate_system.event_log(created_at);