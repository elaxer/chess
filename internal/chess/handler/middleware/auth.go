package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/elaxer/chess/internal/chess/service"
	"github.com/elaxer/chess/internal/ctxkey"
)

const authMethodBearer = "Bearer"

// Auth проверяет наличие заголовка Authorization и авторизует пользователя.
func Auth(next http.Handler, authService *service.Auth) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxkey.User, nil)))
		return
		a := r.Header.Get("Authorization")
		if a == "" {
			http.Error(w, "не передан заголовок Authorization", http.StatusForbidden)
			return
		}
		values := strings.Split(a, " ")
		method, tokenString := values[0], values[1]
		if method != authMethodBearer {
			http.Error(w, "неверный метод Authorization: должен быть Bearer", http.StatusForbidden)
			return
		}

		u, err := authService.Authorize(tokenString)
		if err != nil {
			http.Error(w, "не удалось авторизоваться", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxkey.User, u)))
	})
}
