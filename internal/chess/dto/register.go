package dto

import "github.com/elaxer/chess/internal/chess/model"

// RegisterRequest содержит данные для регистрации.
type RegisterRequest struct {
	Login    string         `json:"login"`
	Password model.Password `json:"password"`
}
