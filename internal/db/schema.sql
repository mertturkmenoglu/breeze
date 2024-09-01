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

CREATE TABLE IF NOT EXISTS sessions (
  id VARCHAR(255) PRIMARY KEY,
  user_id TEXT NOT NULL,
  session_data TEXT,
  created_at TIMESTAMPTZ DEFAULT NOW() NOT NULL,
  expires_at TIMESTAMPTZ NOT NULL
);
