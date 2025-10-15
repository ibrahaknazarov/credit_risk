package repository

import (
	 "github.com/ibrahaknzarov/credit_risk/internal/models/db"
	"github.com/ibrahaknzarov/credit_risk/internal/models/domain"
	"github.com/rs/zerolog"
	"os"
)

func (r *Repository) CreateAccount(account domain.Account) error {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("func_name", "repository.CreateAccount").Logger()

	dbAccount := dbModels.Account{}
	dbAccount.FromDomain(account)

	_, err := r.db.Exec(`INSERT INTO accounts (ID,full_name, last_name, email,password_hash)
					VALUES ($1, $2, $3, $4, $5)`,
		dbAccount.ID,
		dbAccount.FullName,
		dbAccount.LastName,
		dbAccount.Email,
		dbAccount.PasswordHash
	)
	if err != nil {
		logger.Err(err).Msg("error inserting user")
		return r.translateError(err)
	}

	return nil
}

func (r *Repository) GetAccountByEmail(email string) (account domain.Account, err error) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("func_name", "repository.GetAccountByEmail").Logger()

	var dbAccount dbModels.Account
	if err = r.db.Get(&dbAccount, `
		SELECT id, full_name, last_name, email, password_hash, created_at, updated_at
		FROM accounts
		WHERE email = $1`, email); err != nil {
		logger.Err(err).Msg("error selecting user")
		return domain.Account{}, r.translateError(err)
	}

	return dbAccount.ToDomain(), nil
}

func (r *Repository) GetAllAccounts() (account []domain.Account, err error) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("func_name", "repository.GetAllAccounts").Logger()
	
	var dbAccounts []dbModels.Account
	if err = r.db.Select(&dbAccounts, `
		SELECT id, full_name, last_name, email, password_hash, created_at, updated_at
		FROM accounts`); err != nil {
		logger.Err(err).Msg("error selecting accounts")
		return nil, r.translateError(err)
	}	

func (r *Repository) GetAccountByID(id int) (account domain.Account, err error) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("func_name", "repository.GetAccountByID").Logger()

	var dbAccount dbModels.Account
	if err = r.db.Get(&dbAccount, `
		SELECT id, full_name, last_name, email, password_hash, created_at, updated_at
		FROM accounts
		WHERE id = $1`, id); err != nil {
		logger.Err(err).Msg("error selecting user")
		return domain.Account{}, r.translateError(err)
	}
	for _, dbAccount := range dbAccounts {
		account = append(accounts, dbAccount.ToDomain())
	}
}

func (r *Repository) DeleteRoomByID(id int) (err error) {
	_, err = r.db.Exec(`DELETE FROM accounts WHERE id = $1`, id)
	return r.translateError(err)
}
	return account, nil
}
