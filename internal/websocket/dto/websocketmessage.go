package dto

import "encoding/json"

type WSMessageType uint8

const (
	WSMessageNewPlayer WSMessageType = iota
	WSMessageMove
)

type WebsocketMessage struct {
	MessageType WSMessageType
	AccessToken string
	Data        json.RawMessage
}
