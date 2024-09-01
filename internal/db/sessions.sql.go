// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: sessions.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createSession = `-- name: CreateSession :one
INSERT INTO sessions (
  id,
  user_id,
  session_data,
  created_at,
  expires_at
) VALUES (
  $1,
  $2,
  $3,
  $4,
  $5
)
RETURNING id, user_id, session_data, created_at, expires_at
`

type CreateSessionParams struct {
	ID          string
	UserID      string
	SessionData pgtype.Text
	CreatedAt   pgtype.Timestamptz
	ExpiresAt   pgtype.Timestamptz
}

func (q *Queries) CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error) {
	row := q.db.QueryRow(ctx, createSession,
		arg.ID,
		arg.UserID,
		arg.SessionData,
		arg.CreatedAt,
		arg.ExpiresAt,
	)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.SessionData,
		&i.CreatedAt,
		&i.ExpiresAt,
	)
	return i, err
}

const deleteSessionById = `-- name: DeleteSessionById :exec
DELETE FROM sessions
WHERE id = $1 AND user_id = $2
`

type DeleteSessionByIdParams struct {
	ID     string
	UserID string
}

func (q *Queries) DeleteSessionById(ctx context.Context, arg DeleteSessionByIdParams) error {
	_, err := q.db.Exec(ctx, deleteSessionById, arg.ID, arg.UserID)
	return err
}

const getSessionById = `-- name: GetSessionById :one
SELECT sessions.id, sessions.user_id, sessions.session_data, sessions.created_at, sessions.expires_at, users.id, users.email, users.password_hash, users.name, users.role, users.password_reset_token, users.password_reset_expires, users.login_attempts, users.lockout_until, users.last_login, users.created_at, users.updated_at FROM sessions
JOIN users ON users.id = sessions.user_id
WHERE sessions.id = $1 LIMIT 1
`

type GetSessionByIdRow struct {
	Session Session
	User    User
}

func (q *Queries) GetSessionById(ctx context.Context, id string) (GetSessionByIdRow, error) {
	row := q.db.QueryRow(ctx, getSessionById, id)
	var i GetSessionByIdRow
	err := row.Scan(
		&i.Session.ID,
		&i.Session.UserID,
		&i.Session.SessionData,
		&i.Session.CreatedAt,
		&i.Session.ExpiresAt,
		&i.User.ID,
		&i.User.Email,
		&i.User.PasswordHash,
		&i.User.Name,
		&i.User.Role,
		&i.User.PasswordResetToken,
		&i.User.PasswordResetExpires,
		&i.User.LoginAttempts,
		&i.User.LockoutUntil,
		&i.User.LastLogin,
		&i.User.CreatedAt,
		&i.User.UpdatedAt,
	)
	return i, err
}
