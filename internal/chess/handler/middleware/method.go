package middleware

import (
	"net/http"
	"slices"
)

// Method проверяет метод запроса.
// Если метод запроса не соответствует переданным методам, то возвращает ошибку.
func Method(next http.Handler, methods ...string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !slices.Contains(methods, r.Method) {
			http.Error(w, "неверный метод запроса", http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	})

}
