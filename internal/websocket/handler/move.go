package handler

import (
	"encoding/json"
	"errors"

	"github.com/elaxer/chess/pkg/variant/standard/move"
	"github.com/gorilla/websocket"
)

func (h *Websocket) move(conn *websocket.Conn, data json.RawMessage) error {
	move := new(move.Normal)
	if err := json.Unmarshal(data, move); err != nil {
		return err
	}

	session := h.sessionRepository.Find(conn)
	if session == nil {
		return errors.New("игровая сессия не найдена")
	}

	if err := h.gameService.MakeMove(session.Game, move); err != nil {
		return err
	}

	respond(session.PlayerWhite, session.Game.Board)
	respond(session.PlayerBlack, session.Game.Board)

	// todo
	// if move.MoveType.IsGameOver() {
	// 	h.sessionRepository.Remove(session)
	// }

	return nil
}
