package domain
import "time"
type Account struct {
	ID        int       `json:"id" db:"id"`
	UserID    int       `json:"user_id" db:"user_id"`
	Balance   float64   `json:"balance" db:"balance"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

var(
	ErrAccountNotFound = "account not found"
	ErrInsufficientFunds = "insufficient funds"
	ErrInvalidAccountID = "invalid account id"
	ErrAccountCreationFailed = "account creation failed"
	ErrAccountUpdateFailed = "account update failed"
	ErrAccountDeletionFailed = "account deletion failed"
	ErrAccountHistoryNotFound = "account history not found"
	ErrAccountHistoryCreationFailed = "account history creation failed"
	ErrAccountHistoryUpdateFailed = "account history update failed"
	ErrAccountHistoryDeletionFailed = "account history deletion failed"	
)
