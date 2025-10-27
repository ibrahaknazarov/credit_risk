package contracts

import "credit_risk/internal/models/domain"

type RepositoryI interface {
	CreateUser(user domain.User) (err error)
	GetUserByID(id int) (domain.User, error)
	GetUserByUsername(username string) (domain.User, error)

	CreateAccount(account domain.Account) error
	UpdateAccountByID(account domain.Account) error
	GetAllAccounts() (account []domain.Account, err error)
	GetAccountByID(id int) (account domain.Account, err error)
	DeleteAccountByID(id int) (err error)
}
