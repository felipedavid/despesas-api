package middleware

import (
	"context"
	"net/http"
)

func CheckLanguage(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		language := r.Header.Get("Accept-Language")
		r = r.WithContext(context.WithValue(r.Context(), "language", language))
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
