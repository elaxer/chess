package model

import (
	"github.com/elaxer/chess/internal/chess/model"
	"github.com/gorilla/websocket"
)

// Session это игровая сессия между двумя игроками
type Session struct {
	PlayerWhite, PlayerBlack *websocket.Conn
	Game                     *model.Game
}
