package model

import (
	"time"

	"github.com/google/uuid"
)

// BaseModel - это базовая модель для всех моделей.
// Содержит идентификатор и дату создания.
type BaseModel struct {
	ID        uuid.UUID `db:"id"`
	CreatedAt time.Time `db:"created_at"`
}

func newBaseModel() *BaseModel {
	return &BaseModel{uuid.New(), time.Now()}
}
