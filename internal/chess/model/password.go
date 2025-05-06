package model

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
	"golang.org/x/crypto/bcrypt"
)

// Password представляет собой тип пароля пользователя.
type Password string

func (p Password) Validate() error {
	return validation.Validate(
		&p,
		validation.Required,
		validation.Length(6, 32),
		validation.Match(regexp.MustCompile(`\p{Lu}`)),
		validation.Match(regexp.MustCompile(`[\p{P}\p{S}]`)),
	)
}

// Hash хеширует пароль.
func (p Password) Hash() (string, error) {
	ph, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(ph), nil
}

// Compare сравнивает пароль с хешированным паролем.
func (p Password) Compare(hashedPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(p)) == nil
}
