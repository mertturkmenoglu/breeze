CREATE TABLE IF NOT EXISTS users (
    id TEXT PRIMARY KEY,
    email TEXT NOT NULL,
    password_hash TEXT NOT NULL,
    name VARCHAR(128) NOT NULL,
    role VARCHAR(32) DEFAULT 'user' NOT NULL,
    password_reset_token VARCHAR(255),
    password_reset_expires TIMESTAMPTZ,
    login_attempts INT DEFAULT 0,
    lockout_until TIMESTAMPTZ,
    last_login TIMESTAMPTZ DEFAULT NOW() NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT NOW() NOT NULL
);

-- Create indexes
CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);

-- Create triggers
CREATE OR REPLACE TRIGGER update_users_timestamp BEFORE
UPDATE
  ON users FOR EACH ROW EXECUTE FUNCTION update_timestamp();
