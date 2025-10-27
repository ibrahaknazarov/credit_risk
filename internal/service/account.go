package service

import (
	"credit_risk/internal/models/domain"
	
)

func (s *Service) CreateAccount(account domain.Account) error {
	return s.repo.CreateAccount(account)
}
func (s *Service) UpdateAccountByID(account domain.Account) error {
	return s.repo.UpdateAccountByID(account)
}
func (s *Service) DeleteAccountByID(id int) error {
	return s.repo.DeleteAccountByID(id)
}

func (s *Service) GetAllAccounts() (accounts []domain.Account, err error) {
	return s.repo.GetAllAccounts()
}

func (s *Service) GetAccountByEmail(email string) (account domain.Account, err error) {
	return s.repo.GetAccountByEmail(email)
}
