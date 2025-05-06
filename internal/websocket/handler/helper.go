package handler

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

func respond(conn *websocket.Conn, data any) error {
	msg, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		return err
	}

	if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
