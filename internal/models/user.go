package models

import (
	"time"
)

type Role string

const (
	RoleUnknown Role = "unknown"
	RoleUser    Role = "user"
	RoleAdmin   Role = "admin"
)

type User struct {
	ID              uint
	Name            string
	Email           string
	Password        string
	PasswordConfirm string
	Role            string
	CreatedAt       time.Time
	UpdatedAt       *time.Time
}
