package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Role string

const (
	RoleUnknown Role = "unknown"
	RoleUser    Role = "user"
	RoleAdmin   Role = "admin"
)

type User struct {
	ID              int64
	Name            string
	Email           string
	Password        string
	PasswordConfirm string
	Role            Role
	CreatedAt       time.Time
	UpdatedAt       *time.Time
}

type UserClaims struct {
	jwt.StandardClaims
	Email string `json:"email"`
	Role  string `json:"role"`
}
