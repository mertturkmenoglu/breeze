DO $$
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'pagestatus') THEN
    CREATE TYPE PageStatus AS ENUM(
      'NOT_CHECKED',
      'CHECKING',
      'ONLINE',
      'OFFLINE'
    );
  END IF;
END $$;


CREATE TABLE IF NOT EXISTS pages (
  id TEXT PRIMARY KEY,
  name TEXT NOT NULL,
  url TEXT NOT NULL,
  created_at TIMESTAMPTZ DEFAULT NOW() NOT NULL,
  status PageStatus NOT NULL DEFAULT 'NOT_CHECKED',
  uptime INTEGER NOT NULL,
  interval INTEGER NOT NULL,
  last_checked TIMESTAMPTZ DEFAULT NOW() NOT NULL
);
