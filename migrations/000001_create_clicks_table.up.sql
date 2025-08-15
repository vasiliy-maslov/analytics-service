CREATE TABLE IF NOT EXISTS clicks (
    id BIGSERIAL PRIMARY KEY,
    alias VARCHAR(255) NOT NULL,
    "timestamp" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    source VARCHAR(50)
);

CREATE INDEX IF NOT EXISTS idx_clicks_alias ON clicks (alias);
CREATE INDEX IF NOT EXISTS idx_clicks_timestamp ON clicks("timestamp");