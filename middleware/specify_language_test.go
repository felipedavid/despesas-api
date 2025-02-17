package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/felipedavid/saldop/internal/helpers"
	"github.com/felipedavid/saldop/internal/translations"
	"github.com/felipedavid/saldop/test"
)

func TestSpecifyLanguage(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name                string
		languageHeaderValue string
		translator          *translations.Translator
	}{
		{
			"Setting language to portuguese",
			"pt-br",
			translations.PtTranslator,
		},
		{
			"Setting language to english",
			"en-us",
			translations.EnTranslator,
		},
		{
			"Unknown language",
			"whatever",
			translations.EnTranslator,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()

			r, err := http.NewRequest(http.MethodGet, "/", nil)
			if err != nil {
				t.Fatal(err)
			}

			r.Header.Add("Accept-Language", tt.languageHeaderValue)

			next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				wantedTranslator := helpers.GetTranslator(r.Context())
				test.Equal(t, tt.translator, wantedTranslator)
			})

			SpecifyLanguage(next).ServeHTTP(rr, r)
		})
	}
}
