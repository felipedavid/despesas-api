package localizer

import (
	_ "github.com/felipedavid/saldop/translations"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type Localizer struct {
	ID      string
	printer *message.Printer
}

var locales = []Localizer{
	{
		ID:      "en-us",
		printer: message.NewPrinter(language.MustParse("en-us")),
	},
	{
		ID:      "pt-br",
		printer: message.NewPrinter(language.MustParse("pt-br")),
	},
}

func Get(id string) (*Localizer, bool) {
	for _, locale := range locales {
		if id == locale.ID {

			return &locale, true
		}
	}

	return nil, false
}

func (l Localizer) Translate(key message.Reference, args ...interface{}) string {
	return l.printer.Sprintf(key, args...)
}
