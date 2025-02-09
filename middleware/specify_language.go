package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/felipedavid/saldop/internal/helpers"
	"github.com/felipedavid/saldop/internal/translations"
)

func SpecifyLanguage(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var translator *translations.Translator

		language := r.Header.Get("Accept-Language")
		if strings.ToLower(language) == "pt-br" {
			translator = translations.PtTranslator
		} else {
			translator = translations.EnTranslator
		}

		r = r.WithContext(context.WithValue(r.Context(), helpers.TranslatorContextKey, translator))
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
