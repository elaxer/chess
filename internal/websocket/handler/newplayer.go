package handler

import (
	"errors"

	gmodel "github.com/elaxer/chess/internal/chess/model"
	"github.com/elaxer/chess/internal/websocket/model"
	"github.com/gorilla/websocket"
)

func (h *Websocket) newPlayer(u *gmodel.User, conn *websocket.Conn) error {
	if h.sessionRepository.Has(conn) {
		return errors.New("вы уже в игре")
	}

	h.pool.Push(model.ConnUser{Conn: conn, User: u})
	if h.pool.Len() < 2 {
		return nil
	}

	h.mu.Lock()
	playerWhite, playerBlack := h.pool.Pop(), h.pool.Pop()
	h.mu.Unlock()

	game := gmodel.NewGameEmpty(playerWhite.User.ID, playerBlack.User.ID)
	h.sessionRepository.Add(&model.Session{PlayerWhite: playerWhite.Conn, PlayerBlack: playerBlack.Conn, Game: game})

	respond(playerWhite.Conn, game.Board)
	respond(playerBlack.Conn, game.Board)

	return nil
}
