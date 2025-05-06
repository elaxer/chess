package service

import (
	"errors"
	"time"

	"github.com/elaxer/chess/internal/chess/model"
	"github.com/elaxer/chess/internal/chess/repository"
	"github.com/golang-jwt/jwt"
)

// todo: move to config
const secretKey = "secret_key"

var errTokenDecoding = errors.New("ошибка дешифрования токена")

type Auth struct {
	userRepository repository.User
}

func NewAuth(userRepository repository.User) *Auth {
	return &Auth{userRepository}
}

func (s *Auth) Token(sub string) (string, error) {
	payload := jwt.MapClaims{
		"sub": sub,
		"exp": time.Now().Add(time.Second * time.Duration(36000)).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	return token.SignedString([]byte(secretKey))
}

func (s *Auth) Validate(login string, password model.Password) error {
	u, err := s.userRepository.GetByLogin(login)
	if err != nil {
		return err
	}

	if !password.Compare(u.PasswordHash) {
		return errors.New("неверный пароль")
	}

	return nil
}

func (s *Auth) Authenticate(login string, password model.Password) (string, error) {
	if err := s.Validate(login, password); err != nil {
		return "", err
	}

	return s.Token(login)
}

func (s *Auth) Authorize(accessToken string) (*model.User, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("некорректный метод криптографии токена")
		}

		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errTokenDecoding
	}
	if err := claims.Valid(); err != nil {
		return nil, err
	}

	login, ok := claims["sub"]
	if !ok {
		return nil, errTokenDecoding
	}

	return s.userRepository.GetByLogin(login.(string))
}
