package handler

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"strconv"

	"github.com/elaxer/chess/internal/chess/model"
	"github.com/elaxer/chess/internal/chess/repository"
	"github.com/elaxer/chess/internal/ctxkey"
	"github.com/elaxer/chess/internal/database"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// AuthorizedUser возвращает текущего авторизованного пользователя.
func AuthorizedUser(r *http.Request) *model.User {
	u := r.Context().Value(ctxkey.User)
	if u == nil {
		return nil
	}

	return u.(*model.User)
}

func PageUser(w http.ResponseWriter, r *http.Request, userRepository repository.User) (user *model.User, ok bool) {
	id := r.PathValue("id")

	if id == authenticatedUserID {
		return AuthorizedUser(r), true
	}

	uuid, err := uuid.Parse(id)
	if err != nil {
		ResponseError(w, err, http.StatusBadRequest)
		return nil, false
	}

	u, err := userRepository.GetByID(uuid)
	if err != nil {
		ResponseError(w, err, StatusCode(err))
		return nil, false
	}

	return u, true
}

// BeginTx начинает транзакцию и возвращает объект транзакции и контекст.
// Если транзакция не удалась, то запрос завершается ошибкой и метод возвращает false третьим параметром.
func BeginTx(w http.ResponseWriter, db *sqlx.DB) (tx *sqlx.Tx, ctx context.Context, ok bool) {
	tx, ctx, err := database.BeginTx(context.Background(), db)
	if err != nil {
		ResponseError(w, err, http.StatusInternalServerError)
		return nil, nil, false
	}

	return tx, ctx, true
}

// ParamQueryInt возвращает значение параметра из запроса.
// Если параметр не найден или его значение не является числом, то возвращается значение по умолчанию.
func ParamQueryInt(query url.Values, key string, byDefault int) int {
	val := query.Get(key)
	if val == "" {
		return byDefault
	}

	valInt, err := strconv.Atoi(val)
	if err != nil {
		return byDefault
	}

	if valInt < 0 {
		return byDefault
	}

	return valInt
}

// ResponseJSON отправляет ответ в формате JSON.
func ResponseJSON(w http.ResponseWriter, data any, statusCode int) {
	w.Header().Set("Content-type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		ResponseError(w, err, http.StatusInternalServerError)
	}
}

// ReadJSONBody читает тело запроса и декодирует его в структуру.
// Если декодирование не удалось, то запрос завершается ошибкой и метод вторым параметром возвращает false.
func ReadJSONBody[T any](w http.ResponseWriter, r io.Reader) (data T, ok bool) {
	if err := json.NewDecoder(r).Decode(&data); err != nil {
		ResponseError(w, err, http.StatusBadRequest)
		return data, false
	}

	return data, true
}

func ResponseError(w http.ResponseWriter, err error, statusCode int) {
	slog.Error(err.Error())
	http.Error(w, err.Error(), statusCode)
}

// StatusCode возвращает код статуса в зависимости от ошибки.
func StatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	return http.StatusInternalServerError
}
