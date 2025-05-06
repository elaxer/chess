package middleware

import (
	"net/http"

	"github.com/elaxer/chess/internal/chess/handler"
)

// Guest проверяет, авторизирован ли пользователь в системе.
// Если пользователь авторизирован, то он будет перенаправлен на главную страницу.
func Guest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if u := handler.AuthorizedUser(r); u != nil {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}

		next.ServeHTTP(w, r)
	})
}
