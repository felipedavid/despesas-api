package helpers

import (
	"context"

	"github.com/felipedavid/saldop/translations"
)

func GetTranslator(ctx context.Context) *translations.Translator {
	return ctx.Value("translator").(*translations.Translator)
}
