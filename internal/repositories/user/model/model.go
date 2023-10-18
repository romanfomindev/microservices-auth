package model

import (
	"database/sql"
	"time"

	"github.com/romanfomindev/microservices-auth/internal/models"
)

type User struct {
	ID        int64        `db:"id"`
	Name      string       `db:"name"`
	Email     string       `db:"email"`
	Password  string       `db:"password"`
	Role      models.Role  `db:"role"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

type UserCreate struct {
	Name     string      `db:"name"`
	Email    string      `db:"email"`
	Password string      `db:"password"`
	Role     models.Role `db:"role"`
}

type UserUpdate struct {
	Name  string      `db:"name"`
	Email string      `db:"email"`
	Role  models.Role `db:"role"`
}
