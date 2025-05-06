package model

import (
	"github.com/elaxer/chess/internal/chess/model"
	"github.com/gorilla/websocket"
)

// ConnUser это структура, которая хранит вебсокет-соединение и пользователя
type ConnUser struct {
	Conn *websocket.Conn
	User *model.User
}
