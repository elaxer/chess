package handler

import (
	"github.com/elaxer/chess/internal/chess/repository"
)

type User struct {
	userRepository repository.User
}

func NewUser(userRepository repository.User) *User {
	return &User{userRepository}
}
