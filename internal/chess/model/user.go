package model

import validation "github.com/go-ozzo/ozzo-validation"

const (
	// UserInitRating представляет собой начальное значение рейтинга пользователя.
	UserInitRating  = 1500
	UserMinRating   = 0
	UserAddedRating = 50
)

type User struct {
	*BaseModel
	// Login представляет собой логин пользователя.
	Login string `db:"login"`
	// PasswordHash представляет собой хэш пароля пользователя.
	PasswordHash string `db:"password_hash"`
	// IsBot определяет, является ли пользователь ботом.
	IsBot bool `db:"is_bot"`
	// Rating представляет собой рейтинг пользователя.
	Rating int `db:"rating"`
}

func NewUser(login string, password Password) (*User, error) {
	passwordHash, err := password.Hash()
	if err != nil {
		return nil, err
	}

	return &User{
		newBaseModel(),
		login,
		passwordHash,
		false,
		UserInitRating,
	}, nil
}

func (u *User) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Login, validation.Required, validation.Length(3, 28)),
		validation.Field(&u.PasswordHash, validation.Required, validation.Length(0, 255)),
		validation.Field(&u.Rating, validation.Min(UserMinRating)),
	)
}

func (u *User) ChangeRating(rating int) {
	u.Rating += rating
	if u.Rating < UserMinRating {
		u.Rating = UserMinRating
	}
}

func (u *User) String() string {
	return u.Login
}
