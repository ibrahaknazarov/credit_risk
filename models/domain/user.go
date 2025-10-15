package domain
import (
	"time"
)
type User struct {
	ID        int       `json:"id" db:"id"`
	Username  string    `json:"username" db:"username"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"password" db:"password"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt time.Time `json:"deleted_at" db:"deleted_at"`
}

var (
	ErrUserNotFound       = "user not found"
	ErrInvalidUserID      = "invalid user id"
	ErrUserCreationFailed = "user creation failed"
	ErrUserUpdateFailed   = "user update failed"
	ErrUserDeletionFailed = "user deletion failed"
)
