package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/elaxer/chess/internal/chess/model"
	"github.com/elaxer/chess/internal/database"
	"github.com/google/uuid"
	"github.com/huandu/go-sqlbuilder"
)

const usersTableName = "users"

var ErrUserNotFound = fmt.Errorf("%w: пользователь не найден", ErrNotFound)

// User это репозиторий пользователей.
type User interface {
	// GetByID возвращает пользователя по его идентификатору.
	GetByID(id uuid.UUID) (*model.User, error)
	// GetByLogin возвращает пользователя по его логину.
	GetByLogin(login string) (*model.User, error)
	// HasByLogin проверяет существование пользователя с указанным логином.
	// Возвращает true, если пользователь существует, иначе false.
	// Ошибка возвращается, если произошла ошибка при выполнении запроса и не возвращается, если пользователь не найден.
	HasByLogin(login string) (bool, error)
	// Add добавляет пользователя в репозиторий.
	Add(ctx context.Context, u *model.User) error
}

type userDB struct {
	db database.Connection
}

func NewUserDB(db database.Connection) User {
	return &userDB{db}
}

func (r *userDB) GetByID(id uuid.UUID) (*model.User, error) {
	sb := sqlbuilder.NewSelectBuilder()
	query, args := sb.Select("*").From(usersTableName).Where(sb.EQ("id", id)).Limit(1).Build()

	u := new(model.User)
	err := r.db.Get(u, query, args...)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrUserNotFound
	}

	return u, err
}

func (r *userDB) GetByLogin(login string) (*model.User, error) {
	sb := sqlbuilder.NewSelectBuilder()
	query, args := sb.Select("*").From(usersTableName).Where(sb.EQ("login", login)).Limit(1).Build()

	u := new(model.User)
	err := r.db.Get(u, query, args...)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrUserNotFound
	}

	return u, err
}

func (r *userDB) HasByLogin(login string) (bool, error) {
	_, err := r.GetByLogin(login)

	if err != nil && !errors.Is(err, ErrNotFound) {
		return false, err
	}

	return !errors.Is(err, ErrNotFound), nil
}

func (r *userDB) Add(ctx context.Context, u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	query, args := sqlbuilder.
		InsertInto(usersTableName).
		Cols("id", "created_at", "login", "password_hash", "is_bot", "rating").
		Values(u.ID, u.CreatedAt, u.Login, u.PasswordHash, u.IsBot, u.Rating).
		Build()
	_, err := database.TxOrDB(ctx, r.db).ExecContext(ctx, query, args...)

	return err
}
