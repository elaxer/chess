package service

import (
	"context"
	"fmt"

	"github.com/elaxer/chess/internal/chess/model"
	"github.com/elaxer/chess/internal/chess/repository"
)

type User struct {
	userRepository repository.User
}

func NewUser(userRepository repository.User) *User {
	return &User{userRepository}
}

func (s *User) Register(login string, password model.Password) (*model.User, error) {
	if err := password.Validate(); err != nil {
		return nil, err
	}

	if exists, err := s.userRepository.HasByLogin(login); err != nil {
		return nil, err
	} else if exists {
		return nil, fmt.Errorf("%w: пользователь с таким логином уже существует", repository.ErrAlreadyExists)
	}

	u, err := model.NewUser(login, password)
	if err != nil {
		return nil, err
	}

	if err := s.userRepository.Add(context.Background(), u); err != nil {
		return nil, err
	}

	return u, nil
}
