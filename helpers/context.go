package helpers

import (
	"context"
	"net/http"

	"github.com/felipedavid/saldop/models"
	"github.com/felipedavid/saldop/translations"
)

type contextKey string

const userContextKey = contextKey("user")
const TranslatorContextKey = contextKey("translator")

func GetTranslator(ctx context.Context) *translations.Translator {
	return ctx.Value(TranslatorContextKey).(*translations.Translator)
}

func GetUserFromRequestContext(req *http.Request) *models.User {
	user, ok := req.Context().Value(userContextKey).(*models.User)
	if !ok {
		return nil
	}

	return user
}

func SetUserInRequestContext(req *http.Request, user *models.User) *http.Request {
	ctx := context.WithValue(req.Context(), userContextKey, user)
	return req.WithContext(ctx)
}
