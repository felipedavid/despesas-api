package middleware

import (
	"context"
	"net/http"

	"github.com/felipedavid/saldop/translations"
)

func SpecifyLanguage(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var translator *translations.Translator

		language := r.Header.Get("Accept-Language")
		if language == "pt-br" {
			translator = translations.PtTranslator
		} else {
			translator = translations.EnTranslator
		}

		r = r.WithContext(context.WithValue(r.Context(), "translator", translator))
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
