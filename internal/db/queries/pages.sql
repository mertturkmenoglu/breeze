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
WHERE id = $1
ORDER BY pages.created_at DESC
OFFSET $2
LIMIT $3;
