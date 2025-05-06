package repository

import "errors"

var (
	ErrNotFound      = errors.New("не удалось найти в репозитории")
	ErrAlreadyExists = errors.New("уже существует в репозитории")
)
