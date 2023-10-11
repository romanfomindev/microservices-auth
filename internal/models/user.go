package models

import "time"

const ROLE_UNKNOWN = "unknown"
const ROLE_USER = "user"
const ROLE_ADMIN = "admin"

type User struct {
	ID              uint
	Name            string
	Email           string
	Password        string
	PasswordConfirm string
	Role            string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
