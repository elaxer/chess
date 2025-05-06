package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/elaxer/chess/internal/chess/model"
	"github.com/elaxer/chess/internal/database"
	"github.com/google/uuid"
	"github.com/huandu/go-sqlbuilder"
)

const gamesTableName = "games"

// Game это репозиторий шахматных игр.
type Game interface {
	// GetByID возвращает игру по её идентификатору.
	GetByID(id uuid.UUID) (*model.Game, error)
	// GetCount возвращает количество всех сыгранных игр.
	GetCount(ctx context.Context) (int, error)
	// GetLast возвращает последние сыгранные игры.
	GetLast(ctx context.Context, limit, offset int) ([]*model.Game, error)
	// GetCountByUserID возвращает количество всех игр по идентификатору пользователя.
	GetCountByUserID(ctx context.Context, id uuid.UUID) (int, error)
	// GetLastByUserID возвращает последние сыгранные игры по идентификатору пользователя.
	GetLastByUserID(ctx context.Context, userID uuid.UUID, limit, offset int) ([]*model.Game, error)
	// Add добавляет игру в репозиторий.
	Add(ctx context.Context, g *model.Game) error
}

type gameDB struct {
	db database.Connection
}

func NewGameDB(db database.Connection) Game {
	return &gameDB{db}
}

func (r *gameDB) GetByID(id uuid.UUID) (*model.Game, error) {
	b := sqlbuilder.NewSelectBuilder()
	query, args := b.Select("*").From(gamesTableName).Where(b.EQ("id", id)).Limit(1).Build()

	g := new(model.Game)
	err := r.db.Get(g, query, args...)
	if errors.Is(err, sql.ErrNoRows) {
		// todo custom type
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	g.AfterGet()

	return g, nil
}

func (r *gameDB) GetCount(ctx context.Context) (int, error) {
	query, _ := sqlbuilder.Select("COUNT(*)").From(gamesTableName).Build()

	var count int
	err := database.TxOrDB(ctx, r.db).GetContext(ctx, &count, query)

	return count, err
}

func (r *gameDB) GetLast(ctx context.Context, limit, offset int) ([]*model.Game, error) {
	query, args := sqlbuilder.Select("*").
		From(gamesTableName).
		Limit(limit).
		Offset(offset).
		OrderBy("created_at").
		Desc().
		Build()

	games := make([]*model.Game, 0, limit)
	err := database.TxOrDB(ctx, r.db).SelectContext(ctx, &games, query, args...)
	if err != nil {
		return nil, err
	}

	for _, g := range games {
		g.AfterGet()
	}

	return games, nil
}

func (r *gameDB) GetCountByUserID(ctx context.Context, userID uuid.UUID) (int, error) {
	b := sqlbuilder.NewSelectBuilder()

	query, args := b.Select("COUNT(*)").
		From(gamesTableName).
		Where(b.Or(b.Equal("player_white_id", userID), b.Equal("player_black_id", userID))).
		Build()

	var count int
	err := database.TxOrDB(ctx, r.db).GetContext(ctx, &count, query, args...)

	return count, err
}

func (r *gameDB) GetLastByUserID(ctx context.Context, userID uuid.UUID, limit, offset int) ([]*model.Game, error) {
	b := sqlbuilder.NewSelectBuilder()
	query, args := b.Select("*").
		From(gamesTableName).
		Where(b.Or(b.EQ("player_white_id", userID), b.EQ("player_black_id", userID))).
		Limit(limit).
		Offset(offset).
		OrderBy("created_at").
		Desc().
		Build()

	games := make([]*model.Game, 0, limit)
	err := database.TxOrDB(ctx, r.db).SelectContext(ctx, &games, query, args...)
	if err != nil {
		return nil, err
	}

	for _, g := range games {
		g.AfterGet()
	}

	return games, nil
}

func (r *gameDB) Add(ctx context.Context, g *model.Game) error {
	query, args := sqlbuilder.InsertInto(gamesTableName).
		Cols("id", "created_at", "player_white_id", "player_black_id", "moves", "result").
		Values(g.ID, g.CreatedAt, g.PlayerWhiteID, g.PlayerBlackID, g.Moves, g.Result).
		Build()

	_, err := database.TxOrDB(ctx, r.db).ExecContext(ctx, query, args...)

	return err
}
