package db

import (
	"credit_risk/internal/models/domain"
	"time"
)

type Account struct {
	ID					 int       `gorm:"primaryKey;autoIncrement"`
	FirstName        string    `gorm:"type:varchar(100);not null"`
	LastName         string    `gorm:"type:varchar(100);not null"`
	Email            string    `gorm:"type:varchar(100);unique;not null"`
	PasswordHash     string    `gorm:"type:varchar(255);not null"`
	Role             string    `gorm:"type:varchar(50);not null"`
	CreatedAt        time.Time `gorm:"autoCreateTime"`
	UpdatedAt        time.Time `gorm:"autoUpdateTime"`

}

func (r *Account) ToDomain() domain.Account {
	return domain.Account{
		ID:           r.ID,
		FirstName:    r.FirstName,
		LastName:     r.LastName,
		Email:        r.Email,
		PasswordHash: r.PasswordHash,
		Role:         r.Role,
		CreatedAt:    r.CreatedAt,
		UpdatedAt:    r.UpdatedAt,
	}
}

func (r *Account) FromDomain(dr domain.Accout) {
	r.ID = dr.ID
	r.FirstName = dr.FirstName
	r.LastName = dr.LastName
	r.Email = dr.Email
	r.PasswordHash = dr.PasswordHash
	r.Role = dr.Role
	r.CreatedAt = dr.CreatedAt
	r.UpdatedAt = dr.UpdatedAt
}
