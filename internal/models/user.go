package models

import "time"

const (
	ROLE_UNKNOWN = "unknown"
	ROLE_USER    = "user"
	ROLE_ADMIN   = "admin"
)

type User struct {
	ID              uint
	Name            string
	Email           string
	Password        string
	PasswordConfirm string
	Role            string
	CreatedAt       *time.Time
	UpdatedAt       *time.Time
}
