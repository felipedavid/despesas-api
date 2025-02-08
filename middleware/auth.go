package middleware

import (
	"context"
	"log/slog"
	"net/http"
	"strings"

	"github.com/felipedavid/saldop/helpers"
	"github.com/felipedavid/saldop/models"
	"github.com/felipedavid/saldop/storage"
)

func Auth(next http.Handler) http.Handler {
	f := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Vary", "Authorization")

		authHeader := r.Header.Get("Authorization")
		if authHeader != "" {
			splittedHeader := strings.Split(authHeader, " ")
			if len(splittedHeader) != 2 || splittedHeader[0] != "Bearer" {
				http.Error(w, "Malformed authorization header", http.StatusBadRequest)
				return
			}

			token := splittedHeader[1]

			user, err := storage.GetUserByToken(context.Background(), models.TokenScopeAuthentication, token)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusBadRequest)
				slog.Error("get user by token error", "err", err)
				return
			}

			r = helpers.SetUserInRequestContext(r, user)
		}

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(f)
}

func RequireAuthentication(next http.Handler) http.Handler {
	f := func(w http.ResponseWriter, r *http.Request) {
		user := helpers.GetUserFromRequestContext(r)
		if user == nil {
			http.Error(w, "Require authentication", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(f)
}
