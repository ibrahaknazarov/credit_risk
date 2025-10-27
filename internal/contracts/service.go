package contracts

import "github.com/ibrahaknazarov/credit_risk/internal/domain"

type ServiceI interface {
	CreateUser(user domain.User) (err error)
	Authenticate(user domain.User) (int, string, error)

	CreateAccount(account domain.Accounnt) (err error)
	UpdateAccountByID(account domain.Room) (err error)
	DeleteAccountByID(id int) (err error)
	GetAllAccounts() ([]domain.Accounnt, error)
	GetAccountByID(id int) (domain.Accounnt, error)
}
