package service

import (
	"errors"
	"github.com/ibrahaknzarov/credit_risk/internal/errs"
	"github.com/ibrahaknzarov/credit_risk/internal/models/domain"
)

func (s *Service) CreateUser(user domain.User) (err error) {
	// Проверить существует ли пользователь с таким username'ом в бд
	_, err = s.repository.GetUserByUsername(user.Username)
	if err != nil {
		if !errors.Is(err, errs.ErrNotfound) {
			return err
		}
	} else {
		return errs.ErrUsernameAlreadyExists
	}

	// За хэшировать пароль
	user.Password, err = utils.GenerateHash(user.Password)
	if err != nil {
		return err
	}

	user.Role = domain.UserRole

	// Добавить пользователя в бд
	if err = s.repository.CreateUser(user); err != nil {
		return err
	}

	return nil
}

func (s *Service) Authenticate(user domain.User) (int, string, error) {
	// проверить существует ли пользователь с таким username
	userFromDB, err := s.repository.GetUserByUsername(user.Username)
	if err != nil {
		if !errors.Is(err, errs.ErrNotfound) {
			return 0, "", errs.ErrUserNotFound
		}

		return 0, "", err
	}

	// за хэшировать пароль, который получили от пользователя
	user.Password, err = utils.GenerateHash(user.Password)
	if err != nil {
		return 0, "", err
	}

	// проверить правильно ли он указал пароль
	if userFromDB.Password != user.Password {
		return 0, "", errs.ErrIncorrectUsernameOrPassword
	}

	return userFromDB.ID, userFromDB.Role, nil
}
