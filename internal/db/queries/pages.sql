-- name: CreatePage :one
INSERT INTO pages (
  id,
  name,
  url,
  created_at,
  status,
  uptime,
  interval,
  last_checked
) VALUES (
  $1,
  $2,
  $3,
  $4,
  $5,
  $6,
  $7,
  $8
) RETURNING *;

-- name: GetPageById :one
SELECT * FROM pages
WHERE pages.id = $1 LIMIT 1;

-- name: GetPages :many
SELECT * FROM pages
ORDER BY pages.created_at DESC
OFFSET $1
LIMIT $2;

-- name: GetPagesThatNeedChecking :many
SELECT * FROM pages
WHERE pages.status = 'NOT_CHECKED' OR pages.last_checked < NOW() - INTERVAL '1 hour';

-- name: UpdatePageStatus :exec
UPDATE pages
SET status = $1, last_checked = NOW()
WHERE id = $2;
