package repository

import (
	"github.com/ibrahaknzarov/credit_risk/internal/models/db"
	"github.com/ibrahaknzarov/credit_risk/internal/models/domain"
	"github.com/rs/zerolog/log"
	"os"
)

func (r*Repository) CreateUser(user domain.User) (domain.User, error) {
	logger:=zerolog.New(os.Stdout).With().Str("func","CreateUser").Logger()
	_,err = r.db.Exec("INSERT INTO users (id, name, email, password_hash) VALUES ($1, $2, $3, $4)", user.ID, user.Name, user.Email, user.PasswordHash)
	user.FullName,
	uesr.Uesrname,
	user.PasswordHash

	if err != nil {
		logger.Error().Err(err).Msg("Failed to create user")
		return domain.User{}, err
	}
	logger.Info().Msgf("User created with ID: %s", user.ID)
	return user, nil
}

func (r*Repository) GetUserByID(ID int) (domain.User, error) {
	logger:=zerolog.New(os.Stdout).With().Str("func","GetUserByID").Logger()
	var user db.User
	err := r.db.Get(&user, "SELECT id, name, email, password_hash FROM users WHERE id=$1", ID)
	
	if err != nil {
		logger.Error().Err(err).Msgf("Failed to get user with ID: %d", ID)
		return domain.User{}, err
	}
	logger.Info().Msgf("User retrieved with ID: %d", ID)
	return domain.User{
		ID:           user.ID,
		Name:         user.Name,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
	}, nil
}
