package dto

import "github.com/elaxer/chess/internal/chess/model"

// LoginRequest содержит данные для входа.
type LoginRequest struct {
	Login    string         `json:"login"`
	Password model.Password `json:"password"`
}
