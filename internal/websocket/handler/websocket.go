package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/elaxer/chess/internal/chess/service"
	"github.com/elaxer/chess/internal/websocket/dto"
	"github.com/elaxer/chess/internal/websocket/model"
	"github.com/elaxer/chess/internal/websocket/repository"
	"github.com/elaxer/chess/pkg/queue"
	"github.com/gorilla/websocket"
)

type Websocket struct {
	upgrader          websocket.Upgrader
	pool              *queue.Queue[model.ConnUser]
	sessionRepository *repository.Session
	mu                sync.Mutex

	gameService *service.Game
	authService *service.Auth
}

func NewWebsocket(upgrader websocket.Upgrader, gameService *service.Game, authService *service.Auth) *Websocket {
	return &Websocket{
		upgrader,
		queue.New[model.ConnUser](),
		&repository.Session{},
		sync.Mutex{},

		gameService,
		authService,
	}
}

func (h *Websocket) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := h.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	defer conn.Close()

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		if messageType == websocket.CloseMessage {
			h.leave(conn)
			return
		}

		wm := new(dto.WebsocketMessage)
		if err := json.Unmarshal(p, wm); err != nil {
			log.Println(err)
		}

		u, err := h.authService.Authorize(wm.AccessToken)
		if err != nil {
			log.Println(err)
			return
		}

		switch wm.MessageType {
		case dto.WSMessageNewPlayer:
			h.newPlayer(u, conn)
		case dto.WSMessageMove:
			err = h.move(conn, wm.Data)
		}

		if err != nil {
			log.Println(err)
		}
	}
}
