package repository

import (
	"slices"

	"github.com/elaxer/chess/internal/websocket/model"
	"github.com/gorilla/websocket"
)

// Session хранит сессии игры.
type Session struct {
	sessions []*model.Session
}

func (s *Session) Add(session *model.Session) {
	s.sessions = append(s.sessions, session)
}

func (s *Session) Remove(session *model.Session) {
	idx := slices.Index(s.sessions, session)
	s.sessions = slices.Delete(s.sessions, idx, idx)
}

func (s *Session) Find(conn *websocket.Conn) *model.Session {
	for _, session := range s.sessions {
		if session.PlayerWhite == conn || session.PlayerBlack == conn {
			return session
		}
	}

	return nil
}

func (s *Session) Has(conn *websocket.Conn) bool {
	return s.Find(conn) != nil
}
