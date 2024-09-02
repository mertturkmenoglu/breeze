// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"database/sql/driver"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

type Pagestatus string

const (
	PagestatusNOTCHECKED Pagestatus = "NOT_CHECKED"
	PagestatusCHECKING   Pagestatus = "CHECKING"
	PagestatusONLINE     Pagestatus = "ONLINE"
	PagestatusOFFLINE    Pagestatus = "OFFLINE"
)

func (e *Pagestatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Pagestatus(s)
	case string:
		*e = Pagestatus(s)
	default:
		return fmt.Errorf("unsupported scan type for Pagestatus: %T", src)
	}
	return nil
}

type NullPagestatus struct {
	Pagestatus Pagestatus
	Valid      bool // Valid is true if Pagestatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullPagestatus) Scan(value interface{}) error {
	if value == nil {
		ns.Pagestatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Pagestatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullPagestatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Pagestatus), nil
}

type Page struct {
	ID          string
	Name        string
	Url         string
	CreatedAt   pgtype.Timestamptz
	Status      Pagestatus
	Uptime      int32
	Interval    int32
	LastChecked pgtype.Timestamptz
}

type Session struct {
	ID          string
	UserID      string
	SessionData pgtype.Text
	CreatedAt   pgtype.Timestamptz
	ExpiresAt   pgtype.Timestamptz
}

type User struct {
	ID                   string
	Email                string
	PasswordHash         string
	Name                 string
	Role                 string
	PasswordResetToken   pgtype.Text
	PasswordResetExpires pgtype.Timestamptz
	LoginAttempts        pgtype.Int4
	LockoutUntil         pgtype.Timestamptz
	LastLogin            pgtype.Timestamptz
	CreatedAt            pgtype.Timestamptz
	UpdatedAt            pgtype.Timestamptz
}
