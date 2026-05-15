CREATE TABLE urls (
    id          SERIAL PRIMARY KEY,
    original_url TEXT NOT NULL,
    short_code  VARCHAR(20) UNIQUE NOT NULL,
    user_id     INTEGER REFERENCES users(id) ON DELETE CASCADE,
    clicks      INTEGER DEFAULT 0,
    created_at  TIMESTAMP DEFAULT NOW(),
    updated_at  TIMESTAMP DEFAULT NOW()
);