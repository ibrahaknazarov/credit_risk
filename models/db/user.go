package db

import (
	"credit_risk/internal/models/domain"
	"time"
)

type User struct {
	ID           int        `db:"id"`
	FullName     string     `db:"full_name"`
	Username     string     `db:"username"`
	Password     string     `db:"password"`
	Role         string     `db:"role"`
	IsIdentified bool       `db:"is_identified"`
	CreatedAt    time.Time  `db:"created_at"`
	UpdatedAt    *time.Time `db:"updated_at"`	
}

func (u *User) FromDomain(du domain.User) {
    u.ID = du.ID
    u.FullName = du.FullName
    u.Username = du.Username
    u.Password = du.Password
    u.Role = du.Role
    u.IsIdentified = du.IsIdentified
    u.CreatedAt = du.CreatedAt
    u.UpdatedAt = du.UpdatedAt
}
