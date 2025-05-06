package handler

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

// leave обрабатывает событие покидания игры.
// Если игрок покидает игру, то его противнику отправляется сообщение о победе.
// После этого игра завершается и удаляется из сессии.
func (h *Websocket) leave(conn *websocket.Conn) {
	session := h.sessionRepository.Find(conn)
	if session == nil {
		return
	}

	var playerID uuid.UUID
	var victoriousPlayer *websocket.Conn
	if session.PlayerWhite == conn {
		playerID = session.Game.PlayerWhiteID
		victoriousPlayer = session.PlayerBlack
	} else {
		playerID = session.Game.PlayerBlackID
		victoriousPlayer = session.PlayerWhite
	}

	h.gameService.Leave(session.Game, playerID)
	h.sessionRepository.Remove(session)

	respond(victoriousPlayer, "Ваш противник покинул игру")
}
