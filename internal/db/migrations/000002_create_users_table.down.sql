DROP TRIGGER IF EXISTS update_users_timestamp;

DROP INDEX IF EXISTS idx_users_email;

DROP TABLE IF EXISTS users CASCADE;
