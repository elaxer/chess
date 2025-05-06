package handler

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/elaxer/chess/internal/chess/repository"
	"github.com/elaxer/chess/internal/config"
	"github.com/elaxer/chess/pkg/pagination"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Game struct {
	rootDir        string
	db             *sqlx.DB
	userRepository repository.User
	gameRepository repository.Game
}

func NewGame(rootDir string, db *sqlx.DB, userRepository repository.User, gameRepository repository.Game) *Game {
	return &Game{rootDir, db, userRepository, gameRepository}
}

// Game выдает страницу с определенной игрой.
func (h *Game) Game(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		ResponseError(w, err, StatusCode(err))
		return
	}

	game, err := h.gameRepository.GetByID(id)
	if err != nil {
		ResponseError(w, err, StatusCode(err))
		return
	}

	playerWhite, err := h.userRepository.GetByID(game.PlayerWhiteID)
	if err != nil {
		ResponseError(w, err, StatusCode(err))
		return
	}
	playerBlack, err := h.userRepository.GetByID(game.PlayerBlackID)
	if err != nil {
		ResponseError(w, err, StatusCode(err))
		return
	}

	movesJSON, err := json.Marshal(game.Moves)
	if err != nil {
		ResponseError(w, err, http.StatusInternalServerError)
		return
	}

	filenames := []string{
		h.rootDir + config.PublicDir + "/base.html",
		h.rootDir + config.PublicDir + "/game.html",
		h.rootDir + config.PublicDir + "/board.html",
	}

	template.
		Must(template.ParseFiles(filenames...)).
		Execute(w, map[string]any{"Game": game, "PlayerWhite": playerWhite, "PlayerBlack": playerBlack, "MovesJSON": string(movesJSON)})
}

// List выдает страницу со списком всех последних игр.
func (h *Game) List(w http.ResponseWriter, r *http.Request) {
	page := ParamQueryInt(r.URL.Query(), "page", 1)

	tx, ctx, ok := BeginTx(w, h.db)
	if !ok {
		return
	}

	totalCount, err := h.gameRepository.GetCount(ctx)
	if err != nil {
		ResponseError(w, err, StatusCode(err))
		return
	}

	pag := pagination.New(DefaultListLimit, totalCount, page, PaginationButtonsCount)

	games, err := h.gameRepository.GetLast(ctx, pag.Limit(), pag.Offset())
	if err != nil {
		ResponseError(w, err, StatusCode(err))
		return
	}
	for _, g := range games {
		white, err := h.userRepository.GetByID(g.PlayerWhiteID)
		if err != nil {
			ResponseError(w, err, StatusCode(err))
			return
		}
		black, err := h.userRepository.GetByID(g.PlayerBlackID)
		if err != nil {
			ResponseError(w, err, StatusCode(err))
			return
		}

		g.PlayerWhite = white
		g.PlayerBlack = black
	}

	tx.Commit()

	filenames := []string{
		h.rootDir + config.PublicDir + "/base.html",
		h.rootDir + config.PublicDir + "/index.html",
	}

	template.Must(template.ParseFiles(filenames...)).Execute(w, map[string]any{"Games": games})
}

// ListByUser выдает страницу со списком последних игр пользователя.
func (h *Game) ListByUser(w http.ResponseWriter, r *http.Request) {
	u, ok := PageUser(w, r, h.userRepository)
	if !ok {
		return
	}

	page := ParamQueryInt(r.URL.Query(), "page", 1)
	tx, ctx, ok := BeginTx(w, h.db)
	if !ok {
		return
	}

	totalCount, err := h.gameRepository.GetCountByUserID(ctx, u.ID)
	if err != nil {
		ResponseError(w, err, StatusCode(err))
		return
	}

	pag := pagination.New(DefaultListLimit, totalCount, page, PaginationButtonsCount)

	games, err := h.gameRepository.GetLastByUserID(ctx, u.ID, pag.Limit(), pag.Offset())
	if err != nil {
		ResponseError(w, err, StatusCode(err))
		return
	}

	for _, g := range games {
		white, err := h.userRepository.GetByID(g.PlayerWhiteID)
		if err != nil {
			ResponseError(w, err, StatusCode(err))
			return
		}
		black, err := h.userRepository.GetByID(g.PlayerBlackID)
		if err != nil {
			ResponseError(w, err, StatusCode(err))
			return
		}

		g.PlayerWhite = white
		g.PlayerBlack = black
	}

	tx.Commit()

	filenames := []string{
		h.rootDir + config.PublicDir + "/base.html",
		h.rootDir + config.PublicDir + "/games.html",
	}

	template.Must(template.ParseFiles(filenames...)).Execute(w, map[string]any{"Games": games, "User": u, "Pagination": pag})
}
