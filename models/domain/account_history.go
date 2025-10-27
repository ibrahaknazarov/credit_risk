package domain

import "time"

type AccountHistory struct {
	ID        int       `json:"id" db:"id"`
	AccountID int       `json:"account_id" db:"account_id"`
	Change    float64   `json:"change" db:"change"`
	Balance   float64   `json:"balance" db:"balance"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
